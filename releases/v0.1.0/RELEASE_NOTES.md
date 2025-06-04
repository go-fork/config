# Release Notes - v0.1.0

## Overview
Initial release of the go-fork/config package, providing configuration management capabilities for Go applications with dependency injection support.

## What's New
### 🚀 Features
- Configuration management through Manager interface
- ServiceProvider for dependency injection integration
- Support for multiple configuration formats via Viper
- Hot reload capabilities with file watching
- Environment variable support
- Comprehensive test coverage

### 🔧 Core Components
- **Manager Interface**: Core configuration management
- **ServiceProvider**: DI container integration
- **File Watching**: Automatic configuration reload
- **Multi-format Support**: JSON, YAML, TOML, and more

## Dependencies
### Added
- go.fork.vn/di: v0.1.0
- github.com/spf13/viper: v1.20.1
- github.com/fsnotify/fsnotify: v1.9.0

## Installation
```bash
go get go.fork.vn/config@v0.1.0
```

## Contributors
Thanks to the initial development team for making this first release possible.

---
Release Date: 2025-05-XX (Estimated)
