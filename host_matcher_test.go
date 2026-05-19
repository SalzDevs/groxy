package groxy

import "testing"

func TestMatchAllHosts(t *testing.T) {
	matcher := MatchAllHosts()

	for _, host := range []string{"example.com", "example.com:443", ""} {
		if !matcher(host) {
			t.Fatalf("MatchAllHosts()(%q) = false, want true", host)
		}
	}
}

func TestMatchHosts_ExactMatch(t *testing.T) {
	matcher := MatchHosts("example.com")

	if !matcher("example.com") {
		t.Fatal("expected exact host to match")
	}
	if !matcher("EXAMPLE.COM") {
		t.Fatal("expected exact host match to be case-insensitive")
	}
	if matcher("api.example.com") {
		t.Fatal("expected different host not to match")
	}
}

func TestMatchHosts_StripsPort(t *testing.T) {
	matcher := MatchHosts("example.com")

	if !matcher("example.com:443") {
		t.Fatal("expected host with port to match")
	}
}

func TestMatchHosts_WildcardMatch(t *testing.T) {
	matcher := MatchHosts("*.example.com")

	if !matcher("api.example.com") {
		t.Fatal("expected wildcard subdomain to match")
	}
	if !matcher("v1.api.example.com") {
		t.Fatal("expected nested wildcard subdomain to match")
	}
	if matcher("example.com") {
		t.Fatal("expected wildcard not to match root domain")
	}
	if matcher("example.org") {
		t.Fatal("expected different domain not to match")
	}
}

func TestMatchHosts_IgnoresEmptyPatterns(t *testing.T) {
	matcher := MatchHosts("", "  ")

	if matcher("example.com") {
		t.Fatal("expected empty patterns not to match")
	}
}

func TestMatchHosts_IPv6WithPort(t *testing.T) {
	matcher := MatchHosts("2001:db8::1")

	if !matcher("[2001:db8::1]:443") {
		t.Fatal("expected IPv6 host with port to match")
	}
}

func TestMatchHostsPrefix_MatchesPrefix(t *testing.T) {
	matcher := MatchHostsPrefix("internal-")

	for _, host := range []string{"internal-api.example.com", "internal-api.example.com:443", "internal-db"} {
		if !matcher(host) {
			t.Fatalf("MatchHostsPrefix()(%q) = false, want true", host)
		}
	}
}

func TestMatchHostsPrefix_CaseInsensitive(t *testing.T) {
	matcher := MatchHostsPrefix("INTERNAL-")

	if !matcher("internal-api.example.com") {
		t.Fatal("expected case-insensitive prefix match")
	}
}

func TestMatchHostsPrefix_RejectsDifferentPrefix(t *testing.T) {
	matcher := MatchHostsPrefix("internal-")

	for _, host := range []string{"api.example.com", "", "  "} {
		if matcher(host) {
			t.Fatalf("MatchHostsPrefix()(%q) = true, want false", host)
		}
	}
}

func TestMatchHostsSuffix_MatchesSuffix(t *testing.T) {
	matcher := MatchHostsSuffix(".internal")

	for _, host := range []string{"foo.internal", "bar.internal:443", "v1.api.internal"} {
		if !matcher(host) {
			t.Fatalf("MatchHostsSuffix()(%q) = false, want true", host)
		}
	}
}

func TestMatchHostsSuffix_CaseInsensitive(t *testing.T) {
	matcher := MatchHostsSuffix(".INTERNAL")

	if !matcher("foo.internal") {
		t.Fatal("expected case-insensitive suffix match")
	}
}

func TestMatchHostsSuffix_RejectsDifferentSuffix(t *testing.T) {
	matcher := MatchHostsSuffix(".internal")

	for _, host := range []string{"example.com", "", "  "} {
		if matcher(host) {
			t.Fatalf("MatchHostsSuffix()(%q) = true, want false", host)
		}
	}
}

func TestMatchHostsRegex_MatchesPattern(t *testing.T) {
	matcher := MatchHostsRegex(`^api\d+\.example\.com$`)

	for _, host := range []string{"api1.example.com", "api42.example.com:443"} {
		if !matcher(host) {
			t.Fatalf("MatchHostsRegex()(%q) = false, want true", host)
		}
	}
}

func TestMatchHostsRegex_CaseInsensitive(t *testing.T) {
	matcher := MatchHostsRegex(`^api\.example\.com$`)

	if !matcher("API.EXAMPLE.COM") {
		t.Fatal("expected case-insensitive regex match")
	}
}

func TestMatchHostsRegex_RejectsNonMatching(t *testing.T) {
	matcher := MatchHostsRegex(`^api\d+\.example\.com$`)

	for _, host := range []string{"api.example.com", "example.com", "", "  "} {
		if matcher(host) {
			t.Fatalf("MatchHostsRegex()(%q) = true, want false", host)
		}
	}
}

func TestMatchHostsRegex_PanicsOnInvalidPattern(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for invalid regex pattern")
		}
	}()

	_ = MatchHostsRegex(`[invalid`)
}
