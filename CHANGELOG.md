## [0.4.1] - 2026-11-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.4.0

### 💼 Other

- Merge release/v0.4.1 into main
## [0.4.0] - 2026-11-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.9
- *(release)* Update changelog for v0.4.0 (#227)

### 💼 Other

- Merge release/v0.4.0 into main
## [0.3.9] - 2026-11-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.8
- *(release)* Update changelog for v0.3.9 (#223)

### 💼 Other

- Merge release/v0.3.9 into main
## [0.3.8] - 2026-11-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.7
- *(release)* Update changelog for v0.3.8 (#220)

### 💼 Other

- Merge release/v0.3.8 into main
## [0.3.7] - 2026-11-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.6
- *(release)* Update changelog for v0.3.7 (#217)

### 💼 Other

- Merge release/v0.3.7 into main
## [0.3.6] - 2026-11-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.5
- *(ci)* Bump actions/attest-build-provenance from 2 to 4
- *(ci)* Bump anchore/sbom-action from 0.18.0 to 0.24.3
- *(ci)* Bump actions/setup-go from 5 to 7
- *(ci)* Bump actions/checkout from 4 to 7
- *(release)* Update changelog for v0.3.6 (#214)

### 💼 Other

- Merge release/v0.3.6 into main
## [0.3.5] - 2026-11-08

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.4
- *(release)* Update changelog for v0.3.5 (#207)

### 💼 Other

- Merge release/v0.3.5 into main
## [0.3.4] - 2026-11-05

### ⚙️ Miscellaneous Tasks

- *(release)* Fix Releaser complete inputs and use baked-in git-cliff
- *(clean)* Run the cleaner hourly instead of daily
- *(clean)* Lower retention to 5 runs per workflow
- *(sync-crds)* Generate CRDs with mage, not make
- Persist CRDs as a release asset, drop ephemeral artifact uploads
- *(finalize-release)* Put gomarkdoc on PATH after go install
- *(ci)* Update actions
- *(ci)* Update actions
- *(ci)* Update actions
- *(ci)* Update actions
- *(ci)* Update actions
- *(ci)* Update actions
- *(ci)* Update actions
- *(ci)* Update actions
- *(release)* Update changelog for v0.3.4 (#198)
- Sync develop with main after release/v0.3.4
- *(release)* Update changelog for v0.3.4 (#202)
## [0.3.3] - 2026-11-05

### 🐛 Bug Fixes

- Use pre-built developer binary for open-pr.yml
- Use pre-built releaser binary instead of building from source
- Use pre-built releaser binary instead of go run
- Use DEVELOPER_APP_ID and correct repository scope for PR reviewer workflows
- Use developer binary not releaser binary
- Use client-id for create-github-app-token@v3
- Make gh release download conditional on VERSION - planned release flow now builds from source instead of crashing on missing release

### 📚 Documentation

- Standardize workflow header comments
- Standardize open-pr.yml header comments
- Fix open-pr.yml to use pre-built binary
- Standardize workflow header comments
- Standardize apicontract.yml header comments
- Update create-github-app-token to use client-id
- Update create-github-app-token to use client-id

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.2
- *(developer)* Run the developer bot at v0.0.8
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(docs)* Update readme
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- *(ci)* Update bot action versions
- Migrate environments-api bots to v0.1.0 and fix auth
- Fix env-api release workflows to consume env-bots v0.1.0

### 💼 Other

- *(deps)* Bump environments-bots releaser to v0.0.7
- *(deps)* Bump developer binary to v0.1.0
## [0.3.2] - 2026-10-23

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.1
- Run the releaser in the release workflows
- Sync develop with main
- *(release)* Update changelog for v0.3.2 (#98)
## [0.3.1] - 2026-10-22

### 📚 Documentation

- Auto-generate code documentation [skip ci] (#90)

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.3.0
- *(markers)* Add printerColumns
- *(markers)* Add printerColumns, housekeeping and sanitizing
- *(markers)* Add printerColumns, housekeeping and sanitizing
- *(markers)* Add printerColumns, housekeeping and sanitizing
- *(finalize)* Continue when the release pull request is already merged
- *(release)* Update changelog for v0.3.1 (#93)
## [0.3.0] - 2026-10-21

### 🐛 Bug Fixes

- Use the verbatim Apache-2.0 license text

### 📚 Documentation

- Auto-generate code documentation [skip ci]

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.2.9
- *(release)* Make the workflows' commits verified
- Run every job on the self-hosted runners
- *(ci)* Bump actions/setup-go from 6 to 7
- *(release)* Update changelog for v0.2.9
- Do not fail a job when the artifact upload is refused
- Sync develop with main
- *(dependabot)* Open dependency PRs against main
- Sync develop with main
- *(code-docs)* Land generated docs on main through a pull request
- *(release)* Merge the release and its changelog into main through pull requests
- Sync develop with main
- *(release)* Update changelog for v0.3.0 (#87)

### 💼 Other

- Merge release/v0.3.0 into main
## [0.2.9] - 2026-08-20

### 🐛 Bug Fixes

- *(templates)* Move issue templates to the path GitHub actually reads
- *(templates)* Remove Contract Change issue type -- doesn't belong here

### 📚 Documentation

- Add generated code docs and code-docs CI job

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.8
- *(ci)* Bump actions/upload-artifact from 4 to 7
- *(ci)* Bump actions/checkout from 4 to 7
- *(ci)* Bump actions/create-github-app-token from 2 to 3
- *(ci)* Bump softprops/action-gh-release from 2 to 3
- Bring issue/PR template fixes into main

### 💼 Other

- Merge release/v0.2.9 into main
## [0.2.8] - 2026-08-20

### 📚 Documentation

- Remove stale private-repo section from README

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.7
- Sync develop with main after release/v0.2.7
- Add dependabot.yml (gomod, github-actions)
- Bring dependabot.yml into main (Dependabot reads config from default branch)
## [0.2.7] - 2026-08-18

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.2.6
- *(PROJECT)* Lowercase repo/path to match go.mod module path
- Use GitHub App token instead of GH_PAT for release workflows
- Use GitHub App token instead of GH_PAT for release workflows

### 💼 Other

- Bring main into develop (main wins on conflict)
- Merge release/v0.2.7 into main
## [0.2.6] - 2026-07-12

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.5
- Sync develop with main after release/v0.2.5
- *(go.mod)* Use lowercase module path github.com/blanketops/environments-api
## [0.2.5] - 2026-07-12

### 🐛 Bug Fixes

- *(docs)* Correct AGENTS.md API Design section to match actual envelope pattern

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.4
- Sync develop with main after release/v0.2.4
## [0.2.4] - 2026-07-12

### 🐛 Bug Fixes

- *(ci)* Add GOPATH/bin to PATH so mage is found on ubuntu-latest

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.3
- Sync develop with main after release/v0.2.3
- *(docs)* Sync AGENTS.md/PROJECT/README with networks group move and drop BuildTrigger

### 💼 Other

- Merge release/v0.2.4 into main
## [0.2.3] - 2026-07-01

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.2.2
- *(release)* Update changelog for v0.2.2
## [0.2.2] - 2026-07-01

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.1
- Sync develop with main after release/v0.2.1
## [0.2.1] - 2026-07-01

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.2.0
- Sync develop with main after release/v0.2.0
## [0.2.0] - 2026-07-01

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.9
- *(release)* Update changelog for v0.1.9

### 💼 Other

- Merge release/v0.2.0 into main
## [0.1.9] - 2026-07-01

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.8
- *(release)* Update changelog for v0.1.8
## [0.1.8] - 2026-06-12

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.1.7
- Sync develop with main after release/v0.1.7

### 💼 Other

- *(workflows)* Fomart, update breaking magefile
- Merge release/v0.1.8 into main
## [0.1.7] - 2026-06-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.6
- *(release)* Update changelog for v0.1.6

### 💼 Other

- Merge release/v0.1.7 into main
## [0.1.6] - 2026-06-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.5
- *(release)* Update changelog for v0.1.5

### 💼 Other

- *(ci)* # udpate ci release workflows include summary
- *(ci)* # udpate ci release workflows include summary
- *(ci)* # udpate ci release workflows include summary
- Merge release/v0.1.6 into main
## [0.1.5] - 2026-06-11

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.4
- *(release)* Update changelog for v0.1.4

### 💼 Other

- *(ci)* # udpate ci release workflows include summary
- Merge release/v0.1.5 into main
## [0.1.4] - 2026-04-27

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.3
- *(release)* Update changelog for v0.1.3

### 💼 Other

- Merge release/v0.1.4 into main
## [0.1.3] - 2026-04-27

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.2
- *(release)* Update changelog for v0.1.2

### 💼 Other

- Merge release/v0.1.3 into main
## [0.1.2] - 2026-04-27

### ⚙️ Miscellaneous Tasks

- Update apicontract workflow to use mage verify
- Sync develop with main after release/v0.1.1
- Update apicontract workflow to use mage verify
- *(release)* Update changelog for v0.1.1

### 💼 Other

- Merge release/v0.1.2 into main
## [0.1.1] - 2026-04-27

### 🚀 Features

- *(types)* Add githubpayload

### ⚙️ Miscellaneous Tasks

- Sync develop with main after release/v0.1.0
- *(release)* Update changelog for v0.1.0

### 💼 Other

- Merge release/v0.1.1 into main
## [0.1.0] - 2026-04-27

### 🚀 Features

- *(gove)* Add license headers
- *(gov)* Add license to all files

### ⚙️ Miscellaneous Tasks

- *(release)* Update changelog for v0.0.9

### 💼 Other

- *(workflows)* Add release workflows
- *(techdebt)* Remove license header
- Merge release/v0.1.0 into main
## [0.0.9] - 2026-04-05

### 💼 Other

- *(ci)* Run on self-hosted strict, add release to consistent
## [0.0.1] - 2026-02-06
