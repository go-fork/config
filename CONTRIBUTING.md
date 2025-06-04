# Contributing to go-fork/config

Thank you for your interest in contributing to go-fork/config! This document provides guidelines for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Release Process](#release-process)

## Code of Conduct

This project follows the [Go Community Code of Conduct](https://golang.org/conduct). By participating, you are expected to uphold this code.

## Getting Started

### Prerequisites

- Go 1.23 or later
- Git
- Make (optional, for build scripts)

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/config.git
   cd config
   ```
3. Add the upstream remote:
   ```bash
   git remote add upstream https://github.com/go-fork/config.git
   ```

## Development Setup

### Install Dependencies

```bash
go mod download
go mod tidy
```

### Install Development Tools

```bash
# Install mockery for generating mocks
go install github.com/vektra/mockery/v2@latest

# Install golangci-lint for linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Verify Setup

```bash
# Run tests
go test ./...

# Run linter
golangci-lint run

# Generate mocks (if needed)
go generate ./...
```

## Making Changes

### Branching Strategy

- Create feature branches from `main`
- Use descriptive branch names: `feature/add-json-support`, `fix/env-var-parsing`, `docs/update-readme`
- Keep changes focused and atomic

### Coding Standards

1. **Follow Go conventions**: Use `gofmt`, follow naming conventions
2. **Write tests**: All new code should have tests
3. **Document public APIs**: Add godoc comments for public functions/types
4. **Handle errors**: Always handle errors appropriately
5. **Keep interfaces small**: Follow the "interface segregation principle"

### Code Style

```go
// Good: Clear, documented function
// GetConfig retrieves a configuration value by key.
// It returns the value and a boolean indicating if the key exists.
func (m *Manager) GetString(key string) (string, bool) {
    if !m.Has(key) {
        return "", false
    }
    return m.viper.GetString(key), true
}

// Good: Proper error handling
func (m *Manager) ReadInConfig() error {
    if err := m.viper.ReadInConfig(); err != nil {
        return fmt.Errorf("failed to read config: %w", err)
    }
    return nil
}
```

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add support for JSON configuration format
fix: resolve environment variable parsing issue
docs: update usage examples in README
test: add unit tests for SetDefault method
refactor: simplify error handling in ReadInConfig
```

Types:
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation changes
- `test`: Test changes
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `ci`: CI/CD changes

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests with race detection
go test -race ./...

# Run specific tests
go test -run TestManager_GetString ./...
```

### Writing Tests

1. **Test file naming**: `*_test.go`
2. **Test function naming**: `TestFunctionName_Scenario`
3. **Use table-driven tests** for multiple scenarios
4. **Mock external dependencies**

Example test:

```go
func TestManager_GetString(t *testing.T) {
    tests := []struct {
        name     string
        key      string
        setup    func(*Manager)
        want     string
        wantOk   bool
    }{
        {
            name: "existing key",
            key:  "app.name",
            setup: func(m *Manager) {
                m.Set("app.name", "test-app")
            },
            want:   "test-app",
            wantOk: true,
        },
        {
            name:   "non-existing key",
            key:    "missing.key",
            setup:  func(m *Manager) {},
            want:   "",
            wantOk: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            m := NewConfig()
            tt.setup(m)
            
            got, ok := m.GetString(tt.key)
            assert.Equal(t, tt.want, got)
            assert.Equal(t, tt.wantOk, ok)
        })
    }
}
```

### Mocking

Use the provided mock in `mocks/manager.go`:

```go
func TestServiceWithConfig(t *testing.T) {
    mockCfg := &mocks.MockManager{}
    mockCfg.EXPECT().GetString("service.name").Return("test-service", true)
    
    service := NewService(mockCfg)
    assert.Equal(t, "test-service", service.Name())
    
    mockCfg.AssertExpectations(t)
}
```

## Submitting Changes

### Before Submitting

1. **Update your branch**:
   ```bash
   git checkout main
   git pull upstream main
   git checkout your-feature-branch
   git rebase main
   ```

2. **Run quality checks**:
   ```bash
   go test ./...
   golangci-lint run
   go mod tidy
   ```

3. **Update documentation** if needed

### Pull Request Process

1. **Create a pull request** against the `main` branch
2. **Fill out the PR template** completely
3. **Link related issues** using "Fixes #123" or "Closes #123"
4. **Wait for CI checks** to pass
5. **Address review feedback** promptly

### PR Requirements

- [ ] Tests pass
- [ ] Code coverage maintained or improved
- [ ] Documentation updated (if applicable)
- [ ] No breaking changes (unless discussed)
- [ ] Conventional commit messages
- [ ] PR description explains the change

## Release Process

### For Maintainers

1. **Prepare release**:
   ```bash
   # Update version documentation
   ./scripts/archive_release.sh v0.1.3 v0.1.4
   
   # Review and update release notes
   vi releases/v0.1.3/RELEASE_NOTES.md
   ```

2. **Create release**:
   ```bash
   git tag v0.1.3
   git push origin v0.1.3
   ```

3. **Publish**: GitHub Actions will automatically create the release

### Version Guidelines

- **Patch** (`v0.1.3`): Bug fixes, documentation updates
- **Minor** (`v0.2.0`): New features, backward compatible
- **Major** (`v1.0.0`): Breaking changes

## Getting Help

- **Questions**: Use [GitHub Discussions](https://github.com/go-fork/config/discussions)
- **Issues**: Search [existing issues](https://github.com/go-fork/config/issues) first
- **Security**: See [SECURITY.md](SECURITY.md) for security issues

## Recognition

Contributors will be recognized in:
- Release notes
- GitHub contributors list
- Project documentation

Thank you for contributing to go-fork/config!
