# Release Notes - v0.1.1

## Overview
Incremental improvement release focusing on stability, performance optimizations, and enhanced dependency injection integration.

## What's New
### 🐛 Bug Fixes
- Fixed configuration reload edge cases
- Improved error handling in file watching
- Enhanced service provider lifecycle management

### 🔧 Improvements
- Performance optimizations for large configuration files
- Better error messages and logging
- Enhanced test coverage and stability
- Improved documentation and examples

### 📚 Documentation
- Updated API documentation
- Added more usage examples
- Improved README with better getting started guide

## Dependencies
### Updated
- go.fork.vn/di: v0.1.0 → v0.1.1
- Minor dependency updates for security and performance

## Performance
- 15% faster configuration loading for large files
- Reduced memory usage in file watching scenarios
- Optimized service resolution

## Testing
- Increased test coverage to 95%+
- Added integration tests
- Improved mock testing infrastructure

## Installation
```bash
go get go.fork.vn/config@v0.1.1
```

## Migration
Update your dependencies - no code changes required:
```bash
go get go.fork.vn/di@v0.1.1
go get go.fork.vn/config@v0.1.1
go mod tidy
```

---
Release Date: 2025-05-XX (Estimated)
