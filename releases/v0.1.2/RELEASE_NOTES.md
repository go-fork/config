# Release Notes v0.1.2

**Release Date:** June 2, 2025  
**Version:** v0.1.2  
**Module:** go.fork.vn/config  

## Overview

Config v0.1.2 là một bản cập nhật quan trọng nhằm đảm bảo tương thích với go.fork.vn/di v0.1.2, tập trung vào việc nâng cấp Container interface từ concrete type sang interface-based design. Bản cập nhật này cải thiện tính linh hoạt và khả năng test của package thông qua abstraction layer tốt hơn.

## 🚀 Key Improvements

### Container Interface Migration
- **Thay đổi chính**: Migrate từ `*di.Container` sang `di.Container` interface
- **Lợi ích**: Enhanced testability và flexibility thông qua interface-based dependency injection
- **Tương thích ngược**: Hoàn toàn backward compatible - không có breaking changes

### Enhanced Abstraction
- **ServiceProvider**: Cập nhật để sử dụng `di.Container` interface thay vì concrete type
- **Type Safety**: Cải thiện type safety thông qua interface-based design
- **Testability**: Dễ dàng mock và test hơn với interface abstraction

## 📋 What's Changed

### Dependencies
- ⬆️ **Updated**: `go.fork.vn/di` from v0.1.1 to v0.1.2

### Technical Improvements
- 🔧 **Container Interface**: All Container references now use interface instead of concrete type
- 🧪 **Test Infrastructure**: Updated mock imports to use `mocks.MockApplication`
- 📚 **Documentation**: Enhanced documentation to reflect v0.1.2 compatibility
- ✅ **Quality Assurance**: Maintained 99.0% test coverage with updated interface usage

### Testing Enhancements
- **99.0% Test Coverage**: Comprehensive test coverage maintained across all functionalities
- **Mock Updates**: Consistent use of `mocks.MockApplication` for reliable testing
- **Interface Testing**: Enhanced testing capabilities with interface-based dependency injection
- **Static Analysis**: All code passes `go vet` and `golangci-lint` validation

## 🔄 Migration Guide

### For Existing Users
```go
// No changes required - your existing code continues to work
app := di.New()
app.Register(config.NewServiceProvider())

container := app.Container()
cfg := container.MustMake("config").(config.Manager)
```

### Dependency Update
```bash
# Update your go.mod
go get go.fork.vn/di@v0.1.2
go get go.fork.vn/config@v0.1.2
```

## 📊 Quality Metrics

- **Test Coverage**: 99.0% (exceeded 80% requirement)
- **Static Analysis**: ✅ All checks pass
- **Documentation**: ✅ Complete API documentation
- **Performance**: ✅ No performance regressions

## 🧪 Testing

Comprehensive test suite includes:
- Unit tests for all Manager interface methods
- ServiceProvider registration and lifecycle tests
- Mock integration tests with updated interfaces
- Error handling and edge case validation
- Configuration file operations and hot reload functionality

## 📦 Installation

```bash
go get go.fork.vn/config@v0.1.2
```

## 🔗 Dependencies

- **go.fork.vn/di**: v0.1.2 (updated)
- **github.com/spf13/viper**: v1.20.1
- **github.com/fsnotify/fsnotify**: v1.9.0
- **github.com/stretchr/testify**: v1.10.0

## 🏗️ Architecture Benefits

### Interface-Based Design
- Better separation of concerns
- Enhanced testability with easy mocking
- Improved flexibility for future enhancements
- Consistent with Go's interface philosophy

### Dependency Injection
- Seamless integration with updated DI container
- Type-safe service resolution
- Simplified testing with interface mocking
- Better inversion of control

## 🛠️ Development

### Code Quality
- All code passes static analysis (`go vet`, `golangci-lint`)
- Follows Fork Framework coding standards
- Complete godoc documentation
- Idiomatic Go code practices

### Testing Framework
- Table-driven tests for comprehensive coverage
- Interface-based mocking with mockery
- Integration tests with DI container
- Performance and stress testing

## 📝 Notes

- **Backward Compatibility**: 100% backward compatible
- **No Breaking Changes**: All existing APIs remain unchanged
- **Enhanced Testing**: Improved mock generation and testing capabilities
- **Documentation**: Updated to reflect v0.1.2 changes

## 🔮 Next Steps

- Continue monitoring compatibility with future di package updates
- Enhance configuration validation features
- Explore additional configuration sources and formats
- Performance optimizations for large configuration files

---

**Full Changelog**: [v0.1.1...v0.1.2](https://github.com/go-fork/config/compare/v0.1.1...v0.1.2)  
**Documentation**: [go.fork.vn/config](https://pkg.go.dev/go.fork.vn/config)  
**Issues**: [GitHub Issues](https://github.com/go-fork/config/issues)
