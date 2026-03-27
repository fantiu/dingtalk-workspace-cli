# Changelog

All notable changes to this project will be documented in this file.

The format is inspired by Keep a Changelog and this project follows
Semantic Versioning once versioned releases are published.

## [Unreleased]

### Added

- Open-source governance baseline:
  - `LICENSE` (Apache License 2.0)
  - `CODE_OF_CONDUCT.md`
  - `SECURITY.md`
  - `.github/PULL_REQUEST_TEMPLATE.md`
  - `.github/workflows/ci.yml`
  - `.env.example`
  - `docs/architecture.md`
- Open-source policy checks:
  - `scripts/policy/check-open-source-assets.sh`
  - `scripts/policy/open-source-audit.sh`
  - CI integration through `.github/workflows/ci.yml`
- Release hardening:
  - SHA256 checksum generation for packaged artifacts
  - Homebrew tap publish helper

### Changed

- Project license migrated from MIT to Apache License 2.0.
- `README.md` and `CONTRIBUTING.md` expanded for open-source contribution and
  governance guidance.
- `scripts/README.md` now includes open-source policy checks.
