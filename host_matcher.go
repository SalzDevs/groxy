package groxy

import (
	"net"
	"regexp"
	"strings"
)

// MatchAllHosts returns a HostMatcher that matches every host.
func MatchAllHosts() HostMatcher {
	return func(host string) bool {
		return true
	}
}

// MatchHosts returns a HostMatcher for exact and wildcard host patterns.
//
// Patterns are case-insensitive. Hosts may include ports. A pattern beginning
// with "*." matches subdomains only, so "*.example.com" matches
// "api.example.com" but not "example.com".
func MatchHosts(patterns ...string) HostMatcher {
	normalized := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		pattern = normalizeHostPattern(pattern)
		if pattern == "" {
			continue
		}

		normalized = append(normalized, pattern)
	}

	return func(host string) bool {
		host = normalizeHost(host)
		if host == "" {
			return false
		}

		for _, pattern := range normalized {
			if strings.HasPrefix(pattern, "*.") {
				base := strings.TrimPrefix(pattern, "*.")
				if host != base && strings.HasSuffix(host, "."+base) {
					return true
				}
				continue
			}

			if host == pattern {
				return true
			}
		}

		return false
	}
}

// MatchHostsPrefix returns a HostMatcher that matches hosts whose normalized
// form begins with prefix. Matching is case-insensitive.
func MatchHostsPrefix(prefix string) HostMatcher {
	prefix = strings.ToLower(strings.TrimSpace(prefix))
	return func(host string) bool {
		host = normalizeHost(host)
		return host != "" && strings.HasPrefix(host, prefix)
	}
}

// MatchHostsSuffix returns a HostMatcher that matches hosts whose normalized
// form ends with suffix. Matching is case-insensitive.
func MatchHostsSuffix(suffix string) HostMatcher {
	suffix = strings.ToLower(strings.TrimSpace(suffix))
	return func(host string) bool {
		host = normalizeHost(host)
		return host != "" && strings.HasSuffix(host, suffix)
	}
}

// MatchHostsRegex returns a HostMatcher that matches hosts against a regular
// expression. Matching is case-insensitive. The pattern is compiled with
// regexp.MustCompile, so an invalid pattern panics at registration time.
func MatchHostsRegex(pattern string) HostMatcher {
	re := regexp.MustCompile(pattern)
	return func(host string) bool {
		host = normalizeHost(host)
		return host != "" && re.MatchString(host)
	}
}

func normalizeHostPattern(pattern string) string {
	return normalizeHost(pattern)
}

func normalizeHost(host string) string {
	host = strings.TrimSpace(strings.ToLower(host))
	if host == "" {
		return ""
	}

	if strings.HasPrefix(host, "[") {
		if withoutPort, _, err := net.SplitHostPort(host); err == nil {
			return strings.Trim(withoutPort, "[]")
		}
		return strings.Trim(host, "[]")
	}

	if withoutPort, _, err := net.SplitHostPort(host); err == nil {
		return withoutPort
	}

	if strings.Count(host, ":") == 1 {
		name, port, ok := strings.Cut(host, ":")
		if ok && port != "" {
			return name
		}
	}

	return strings.TrimSuffix(host, ".")
}
