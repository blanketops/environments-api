# Security Policy

## Supported Versions

BlanketOps Environments API is pre-1.0 — all API groups are currently
`v1alpha1` and may change in backward-incompatible ways. Only the **latest
released tag** receives security fixes; there's no long-term support branch
yet.

| Version        | Supported          |
| -------------- | ------------------- |
| Latest release | :white_check_mark:  |
| Older releases | :x:                 |

This will change once an API version stabilizes to `v1beta1`/`v1` with a
defined support window.

## Scope

This repository defines Kubernetes API contracts (CRDs and Go types) — it
contains no controllers, runtimes, or deployed services of its own. Relevant
security concerns here are things like malformed/malicious CRD manifests,
supply-chain integrity of the published CRD/OpenAPI bundle, and vulnerable
dependencies in `go.mod`. Reconciliation and deployment-time concerns belong
to downstream operator repos.

## Reporting a Vulnerability

Please report security vulnerabilities privately — **do not open a public
GitHub issue**.

Use [GitHub's private vulnerability reporting](https://github.com/blanketops/environments-api/security/advisories/new)
for this repository. If that's unavailable, contact a maintainer directly
through GitHub.

Include what you'd include in any good report: affected version/commit,
reproduction steps, and impact. We'll acknowledge new reports within a few
business days. This is a small, actively-developed project — there's no
formal SLA yet, but valid reports are prioritized over other work.

Once a fix is available, we'll coordinate disclosure timing with you and
credit you in the release notes unless you'd prefer otherwise.
