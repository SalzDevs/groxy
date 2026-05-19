# Roadmap

Groxy is pre-v1. This roadmap tracks what shipped and what's planned next.

## Done

### v0.3.0 — HTTPS inspection

- Explicit opt-in HTTPS inspection with local CA.
- `HostMatcher`, `MatchHosts`, `MatchAllHosts`.
- `CAConfig`, `NewCA`, `LoadCAFiles`, `CA.WriteFiles`.
- Per-host certificate generation and cache.
- Certificate renewal before expiry.
- Fail-closed defaults with `PassthroughOnError`.
- Request/response hooks on intercepted HTTPS.
- Body transforms on intercepted HTTPS.
- `BlockHost` on intercepted HTTPS.

### v0.3.1 — Documentation improvements

- Quickstart and forward proxy building guide.
- CA trust instructions.
- Open-source hygiene files.

### v0.4.0 — Access logs

- `AccessLog` middleware for one-line HTTP and CONNECT traffic logs.
- Access log example.
- No credential or header leakage in logs.

### v0.4.1 — Timeout documentation

- Timeout semantics guide covering client-to-proxy and proxy-to-upstream behavior.
- Clarified defaults, zero values, and validation.

### v0.5.0 — Proxy authentication

- `ProxyBasicAuth` for static HTTP Basic proxy authentication.
- `ProxyBasicAuthFunc` for custom validators.
- Protects both HTTP proxy requests and HTTPS CONNECT tunnels.
- `407 Proxy Authentication Required` response handling.
- Credential stripping before upstream.
- No reauthentication for inspected HTTPS requests.
- Proxy auth parsing helpers.
- Proxy auth guide and runnable example.

### v0.5.1 — Custom block and error response examples

- Custom block/error response guide.
- Runnable custom block response example.
- README and docs links updated.

### v0.5.2 — HTTPS inspection benchmarks

- Certificate cache miss and hit benchmarks.
- Intercepted HTTPS forwarding benchmarks.
- Benchmarks with hooks and response body transforms.

## Next

### v0.6.0 — HTTPS inspection hardening

Goal: improve safety, configurability, and production ergonomics for HTTPS
inspection.

Potential work:

- richer host matching controls (`MatchHostsPrefix`, `MatchHostsSuffix`, `MatchHostsRegex`)
- custom upstream TLS settings
- better certificate lifecycle visibility
- optional certificate persistence hooks
- more documentation for browser and OS trust setup

### v1.0.0 — API stabilization

Goal: stabilize the public API after real-world feedback.

Before v1:

- review exported names and docs
- confirm middleware API ergonomics
- document compatibility guarantees
- finalize error handling behavior
- ensure examples cover common use cases

## Good first issue ideas

- Add richer host matching controls.
- Add more docs for installing the Groxy CA in common browsers and OSes.
- Add a streaming request/response body transform API design discussion.
