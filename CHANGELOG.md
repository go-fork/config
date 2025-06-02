# Changelog

## [Unreleased]

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

[Unreleased]: https://github.com/go-fork/config/compare/v0.1.1...HEAD
[v0.1.1]: https://github.com/go-fork/config/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/go-fork/config/releases/tag/v0.1.0
