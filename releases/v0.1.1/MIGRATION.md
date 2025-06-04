# Migration Guide - v0.1.1

## Overview
This guide helps you migrate from v0.1.0 to v0.1.1. This release is fully backward compatible with performance and stability improvements.

## Prerequisites
- Go 1.21 or later
- go.fork.vn/config v0.1.0

## Quick Migration Checklist
- [x] Update dependencies
- [x] Run tests to ensure compatibility
- [x] No code changes required

## Breaking Changes
**None** - This release maintains full backward compatibility.

## Step-by-Step Migration

### Step 1: Update Dependencies
```bash
go get go.fork.vn/di@v0.1.1
go get go.fork.vn/config@v0.1.1
go mod tidy
```

### Step 2: Run Tests
```bash
go test ./...
```

### Step 3: Enjoy Performance Improvements
Your existing code will automatically benefit from:
- 15% faster configuration loading
- Reduced memory usage
- Better error handling
- Enhanced stability

## What Changed
### Performance Improvements
- Optimized configuration loading for large files
- Reduced memory footprint in file watching
- Faster service resolution

### Bug Fixes
- Fixed configuration reload edge cases
- Improved file watching reliability
- Better error handling and recovery

### Enhanced Testing
- Increased test coverage to 95%+
- Better integration test coverage
- Improved mock testing infrastructure

## Common Issues and Solutions
This is a stable release, but if you encounter issues:

### Issue: Performance Regression
**Unlikely Problem**: Slower performance after update  
**Solution**: 
```bash
go clean -modcache
go mod download
go build -a ./...
```

### Issue: File Watching Issues
**Solution**: The new version has improved file watching - restart your application to benefit from the fixes.

## Getting Help
- Check the [documentation](https://pkg.go.dev/go.fork.vn/config@v0.1.1)
- Search [existing issues](https://github.com/go-fork/config/issues)
- Create a [new issue](https://github.com/go-fork/config/issues/new) if needed

## Rollback Instructions
If needed (though unlikely):

```bash
go get go.fork.vn/config@v0.1.0
go get go.fork.vn/di@v0.1.0
go mod tidy
```

## Benefits of Upgrading
- 15% performance improvement in configuration loading
- Better stability and error handling
- Reduced memory usage
- Enhanced developer experience
- Foundation for future interface improvements

---
**Need Help?** This upgrade should be seamless and provide immediate performance benefits.
