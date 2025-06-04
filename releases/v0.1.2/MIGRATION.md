# Migration Guide - v0.1.2

## Overview
This guide helps you migrate from v0.1.1 to v0.1.2. The good news is that **no code changes are required** - this release is 100% backward compatible.

## Prerequisites
- Go 1.23 or later
- go.fork.vn/config v0.1.1 or earlier

## Quick Migration Checklist
- [x] Update dependencies (only step required)
- [x] Run tests to ensure compatibility
- [x] No code changes needed

## Breaking Changes
**None** - This release maintains full backward compatibility.

## Step-by-Step Migration

### Step 1: Update Dependencies
The only change required is updating your dependencies:

```bash
go get go.fork.vn/di@v0.1.2
go get go.fork.vn/config@v0.1.2
go mod tidy
```

### Step 2: Verify Compatibility
Run your tests to ensure everything works as expected:

```bash
go test ./...
```

### Step 3: Optional - Leverage New Interface Features
While not required, you can now take advantage of enhanced testability:

```go
// Your existing code continues to work unchanged:
app := di.New()
app.Register(config.NewServiceProvider())

container := app.Container()
cfg := container.MustMake("config").(config.Manager)

// Enhanced testing capabilities are now available
// (see documentation for advanced testing patterns)
```

## What Changed Internally
While your code doesn't need to change, here's what improved under the hood:

### Container Interface
- **Before**: Used concrete `*di.Container` type
- **After**: Uses `di.Container` interface
- **Impact**: Better testability and flexibility, no API changes

### Enhanced Testing
- Improved mock generation capabilities
- Better interface-based testing patterns
- More flexible dependency injection patterns

## Common Issues and Solutions
Since this is a fully compatible release, you shouldn't encounter any issues. However:

### Issue: Build Errors After Update
**Unlikely Problem**: Build fails after dependency update  
**Solution**: 
```bash
go clean -modcache
go mod download
go mod tidy
```

### Issue: Test Failures
**Unlikely Problem**: Existing tests fail  
**Solution**: This shouldn't happen, but if it does:
1. Check if you're using internal/private APIs (not recommended)
2. Ensure all dependencies are updated to compatible versions

## Getting Help
- Check the [documentation](https://pkg.go.dev/go.fork.vn/config@v0.1.2)
- Search [existing issues](https://github.com/go-fork/config/issues)
- Create a [new issue](https://github.com/go-fork/config/issues/new) if needed

## Rollback Instructions
If you need to rollback (though it shouldn't be necessary):

```bash
go get go.fork.vn/config@v0.1.1
go get go.fork.vn/di@v0.1.1
go mod tidy
```

## Benefits of Upgrading
- Enhanced testability through interface abstraction
- Better compatibility with future DI container updates
- Improved development experience
- Foundation for future feature enhancements

---
**Need Help?** This migration should be seamless. If you encounter any issues, please open an issue on GitHub.
