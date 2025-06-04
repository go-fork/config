# Release Notes - v0.1.3

## Overview
This release restructures the project to match go-fork/di standards with comprehensive CI/CD pipeline, release management system, and enhanced testing infrastructure.

## What's New
### 🚀 Features
- Complete CI/CD pipeline with GitHub Actions workflows
- Automated release management system with templates and scripts
- Comprehensive test coverage (98.0%)
- Automated dependency updates via Dependabot

### 🐛 Bug Fixes
- Fixed ServiceProvider.Boot panic behavior test assertions
- Corrected import paths for di_mocks package compatibility

### 🔧 Improvements
- Project structure now matches go-fork/di standards
- Updated Go version requirement to 1.23+ for Fork framework compatibility
- Enhanced code quality with golangci-lint integration
- Race condition detection in tests

### 📚 Documentation
- Added CONTRIBUTING.md with comprehensive contribution guidelines
- Added SECURITY.md with security policy and vulnerability reporting
- Complete release workflow documentation in releases/README.md
- Scripts documentation for automated release management

## Breaking Changes
### ⚠️ Important Notes
- Breaking change 1 (if any)
- Breaking change 2 (if any)

## Migration Guide
See [MIGRATION.md](./MIGRATION.md) for detailed migration instructions.

## Dependencies
### Updated
- dependency-name: vX.Y.Z → vA.B.C

### Added
- new-dependency: vX.Y.Z

### Removed
- removed-dependency: vX.Y.Z

## Performance
- Benchmark improvement: X% faster in scenario Y
- Memory usage: X% reduction in scenario Z

## Security
- Security fix for vulnerability X
- Updated dependencies with security patches

## Testing
- Added X new test cases
- Improved test coverage to X%

## Contributors
Thanks to all contributors who made this release possible:
- @contributor1
- @contributor2

## Download
- Source code: [go.fork.vn/config@v0.1.3]
- Documentation: [pkg.go.dev/go.fork.vn/config@v0.1.3]

---
Release Date: 2025-06-04
