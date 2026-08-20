## What

<!-- One or two sentences. What does this PR change? -->

## Why

<!-- The intent. Link the issue if one exists: Closes #123 -->

## Domain

<!-- Mark all that apply -->

* [ ] `environments`
* [ ] `events`
* [ ] `sources`
* [ ] `networks`
* [ ] `common`
* [ ] CI / tooling / docs

## API impact

* [ ] No API surface change
* [ ] `v1alpha1` — free to change
* [ ] `v1beta1` — backwards-compatible only, deprecations allowed
* [ ] `v1` — **breaking change** (requires version bump + changelog entry + migration note)

<!-- If breaking: what breaks, and what must consumers do? -->

## Checklist

* [ ] `mage verify` passes locally (fmt, vet, deepcopy generate, CRD manifests)
* [ ] Generated deepcopy/CRD manifests committed and in sync with the Go types
* [ ] No reconciliation logic added — this repo owns schemas and types only
* [ ] BlanketOps labels present where required (`environments.blanketops.dev/*`)
* [ ] Commit messages follow Conventional Commits

## Notes for reviewer

<!-- Anything non-obvious: design trade-offs, deferred follow-ups, areas needing close attention -->
