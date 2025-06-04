# Changelog

## [Unreleased]

## v0.1.3 - 2025-06-04

### Added
- **GitHub Actions CI/CD Pipeline**: Complete automation for testing, building, and releasing
  - Continuous Integration workflow with Go 1.23-1.24 matrix testing
  - Automated release workflow triggered by version tags
  - Weekly dependency update automation
- **Release Management System**: Structured release workflow with templates
  - Release documentation templates (RELEASE_NOTES.md, MIGRATION.md, RELEASE_SUMMARY.md)
  - Versioned release archives in `releases/` directory
  - Automated scripts for release template generation and archiving
- **GitHub Configuration**: Professional GitHub repository setup
  - Issue and PR templates for standardized contributions
  - CODEOWNERS file for code ownership management
  - Dependabot configuration for automated dependency updates
  - Funding configuration for project sponsorship
- **Documentation Enhancement**: Comprehensive contribution and security documentation
  - CONTRIBUTING.md with detailed contribution guidelines
  - SECURITY.md with security policy and vulnerability reporting procedures
  - Enhanced README.md with updated Go version requirements

### Changed
- **Go Version Requirement**: Updated from Go 1.21+ to Go 1.23+ for Fork framework compatibility
- **Dependency Update**: Updated `go.fork.vn/di` from v0.1.2 to v0.1.3
- **Project Structure**: Restructured to match go-fork/di project standards
  - Added `.github/` folder with workflows and templates
  - Added `releases/` folder for version management
  - Added `scripts/` folder for automation tools

### Fixed
- **Test Assertions**: Corrected ServiceProvider.Boot panic behavior test expectations
- **Import Compatibility**: Fixed import paths for di_mocks package in test files

### Technical Improvements
- **Code Quality**: Enhanced linting and code quality checks
  - golangci-lint integration with zero linting errors
  - Race condition detection in test suite
  - 98.0% test coverage maintained
- **Automation Scripts**: Created executable scripts for release management
  - `scripts/create_release_templates.sh`: Generate release templates
  - `scripts/archive_release.sh`: Archive and manage releases
- **CI/CD Integration**: Multi-platform testing and automated workflows
  - Matrix testing across Go 1.23 and 1.24
  - Automated dependency updates via Dependabot
  - Release automation with GitHub Actions

### Migration Notes
- Update Go version to 1.23+ to ensure compatibility
- No breaking changes to public API - all existing code remains compatible
- Enhanced project structure for better maintainability and contribution

## v0.1.2 - 2025-06-02

### Changed
- **Dependency Update**: Updated `go.fork.vn/di` from v0.1.1 to v0.1.2
  - **Container Interface Migration**: Changed from `*di.Container` to `di.Container` interface
  - Enhanced interface-based dependency injection for better testability
  - Improved abstraction layer for container operations

### Technical Improvements
- **Container Interface**: Updated all Container references to use interface instead of concrete type
  - ServiceProvider registration now uses `di.Container` interface
  - Improved flexibility and testability through interface-based design
- **Test Infrastructure**: Updated mock imports to use `mocks.MockApplication`
- **Documentation**: Enhanced documentation to reflect v0.1.2 compatibility
- **Code Quality**: Maintained 99.0% test coverage with updated interface usage

### Testing
- **99.0% Test Coverage**: Comprehensive test coverage maintained
- **Mock Updates**: Updated to use `mocks.MockApplication` for consistency
- **Interface Testing**: Enhanced testing with interface-based dependency injection
- **Static Analysis**: All code passes `go vet` and `golangci-lint` checks

### Migration Notes
- Applications using config package should update `go.fork.vn/di` to v0.1.2
- No breaking changes to public API - all existing code remains compatible
- Enhanced type safety through interface-based container usage

## v0.1.1 - 2025-06-02

### Changed
- **Dependency Update**: Updated `go.fork.vn/di` from v0.1.0 to v0.1.1
  - Enhanced type safety with Application interface improvements
  - Improved documentation with strongly typed method signatures
  - Better consistency across ServiceProvider interfaces

### Technical Improvements
- **ServiceProvider Interface**: Updated to use strongly typed `di.Application` instead of `interface{}`
  - `Register(app di.Application)` - Type-safe application parameter
  - `Boot(app di.Application)` - Consistent interface usage
- **Test Coverage**: Maintained 100% test coverage with updated mocking strategy
- **Mock Integration**: Now uses `go.fork.vn/di/mocks` for consistent testing

### Testing
- **100% Test Coverage**: All functions and statements covered
- **25 Test Functions**: Comprehensive test suite with 65+ assertions
- **Mock Support**: Full integration with di package mocks
- **Coverage Reports**: HTML and text coverage reports generated

## v0.1.0 - 2025-05-31

### Added
- **Configuration Management**: Comprehensive configuration management system for Go applications
- **Multiple Sources**: Support for YAML, JSON, TOML, and environment variables
- **File Watching**: Automatic configuration reload on file changes with fsnotify
- **Type Safety**: Strongly typed configuration with struct binding
- **Environment Override**: Environment variable support with prefix configuration
- **Validation**: Built-in validation for required fields and value constraints
- **DI Integration**: Seamless integration with Dependency Injection container
- **Hot Reload**: Live configuration updates without application restart
- **Default Values**: Support for default values and fallback configurations
- **Nested Configuration**: Support for complex nested configuration structures
- **Path Resolution**: Intelligent configuration file path resolution
- **Error Handling**: Comprehensive error handling for configuration loading
- **Testing Support**: Mock implementations for testing scenarios
- **Viper Integration**: Advanced configuration management with Viper
- **Type-safe Access**: Comprehensive API for type-safe configuration access
  - String, Int, Bool, Float value retrieval
  - Duration, Time value retrieval
  - Slice, Map, and complex structure support
- **File Search**: Automatic configuration file discovery with search paths
- **Configuration Merging**: Support for configuration mounting and merging
- **Mock Support**: Full compatibility with testify/mock framework
- **Performance**: Optimized for reading large configuration files
- **Security**: Support for secure credential storage

### Technical Details
- Initial release as standalone module `go.fork.vn/config`
- Repository located at `github.com/go-fork/config`
- Built with Go 1.23.9
- Full test coverage and documentation included
- Integration with Viper for advanced configuration features
- Thread-safe configuration management
- Easy mock regeneration with `mockery --name Manager` command

### Dependencies
- `github.com/fsnotify/fsnotify`: File system notifications for hot reload
- `github.com/spf13/viper`: Advanced configuration management
- `go.fork.vn/di`: Dependency injection integration

[Unreleased]: https://github.com/go-fork/config/compare/v0.1.3...HEAD
[v0.1.3]: https://github.com/go-fork/config/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/go-fork/config/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/go-fork/config/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/go-fork/config/releases/tag/v0.1.0
