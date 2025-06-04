# Migration Guide - v0.1.3

## Overview
This guide helps you migrate from v0.1.2 to v0.1.3. This is primarily a project infrastructure release with minimal API changes.

## Prerequisites
- Go 1.23 or later (updated from Go 1.21+)
- Previous version installed

## Quick Migration Checklist
- [ ] Update Go version to 1.23+ if needed
- [ ] Update dependencies with `go mod tidy`
- [ ] Verify tests still pass with `go test ./...`
- [ ] No code changes required for most users

## Breaking Changes

### Go Version Requirement
```go
// Previous requirement
go 1.21

// New requirement (v0.1.3)
go 1.23
```

#### No API Changes
This release focuses on project infrastructure improvements. All existing APIs remain unchanged and fully compatible.

#### Changed Types
```go
// Old type definition
type OldConfig struct {
    Field1 string
    Field2 int
}

// New type definition
type NewConfig struct {
    Field1 string
    Field2 int64 // Changed from int
    Field3 bool  // New field
}
```

### Configuration Changes
If you're using configuration files:

```yaml
# Old configuration format
old_setting: value
deprecated_option: true

# New configuration format
new_setting: value
# deprecated_option removed
new_option: false
```

## Step-by-Step Migration

### Step 1: Update Dependencies
```bash
go get go.fork.vn/config@v0.1.3
go mod tidy
```

### Step 2: Update Import Statements
```go
// If import paths changed
import (
    "go.fork.vn/config" // Updated import
)
```

### Step 3: Update Code
Replace deprecated function calls:

```go
// Before
result := config.OldFunction(param)

// After
result := config.NewFunction(param, defaultValue)
```

### Step 4: Update Configuration
Update your configuration files according to the new schema.

### Step 5: Run Tests
```bash
go test ./...
```

## Common Issues and Solutions

### Issue 1: Function Not Found
**Problem**: `undefined: config.OldFunction`  
**Solution**: Replace with `config.NewFunction`

### Issue 2: Type Mismatch
**Problem**: `cannot use int as int64`  
**Solution**: Cast the value or update variable type

## Getting Help
- Check the [documentation](https://pkg.go.dev/go.fork.vn/config@v0.1.3)
- Search [existing issues](https://github.com/go-fork/config/issues)
- Create a [new issue](https://github.com/go-fork/config/issues/new) if needed

## Rollback Instructions
If you need to rollback:

```bash
go get go.fork.vn/config@previous-version
go mod tidy
```

Replace `previous-version` with your previous version tag.

---
**Need Help?** Feel free to open an issue or discussion on GitHub.
