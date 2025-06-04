# Config Package Overview

## Architecture

The config package is built around the Fork Framework's core principles of dependency injection, service providers, and interface-based design. It provides a unified configuration management system that integrates seamlessly with the framework's ecosystem.

### Core Components

#### Manager Interface
The central `config.Manager` interface provides all configuration operations:
- Value retrieval with type safety
- Configuration file management
- Environment variable integration
- Live configuration watching

#### Service Provider
The `ServiceProvider` registers the config manager as a singleton in the DI container, making it available throughout the application via `app.MustMake("config")`.

#### Type System
The package supports all Go primitive types plus:
- Slices and arrays
- Maps with various key/value combinations
- Time and duration types
- Custom struct unmarshaling

## Design Principles

### 1. Framework Integration First
The config package is designed specifically for Fork Framework applications. All examples and patterns assume usage within the framework's DI container system.

### 2. Type Safety
All value retrieval methods return both the value and an existence flag, preventing runtime panics from missing configuration keys.

```go
if value, exists := configManager.GetString("key"); exists {
    // Safe to use value
}
```

### 3. Environment Precedence
Environment variables automatically override file-based configuration when using the configured prefix (default: `FORK_`).

### 4. Fail-Safe Defaults
The system encourages setting reasonable defaults and graceful degradation when configuration is missing.

## Configuration Sources

The config manager supports multiple configuration sources in order of precedence:

1. **Explicit Set Values** - Values set programmatically via `Set()`
2. **Environment Variables** - With configured prefix (e.g., `FORK_DATABASE_HOST`)
3. **Configuration Files** - YAML, JSON, TOML, or other supported formats
4. **Default Values** - Set via `SetDefault()`

### Configuration File Resolution

The manager searches for configuration files in the following order:
1. Exact file path (if specified with `SetConfigFile()`)
2. Config name + type in search paths (e.g., `app.yaml`)
3. Environment-specific overlays (e.g., `app.production.yaml`)

## Framework Integration Patterns

### Service Provider Registration
```go
// Automatic registration in Fork Framework
app := app.NewApplication()
// Config provider is auto-registered

// Access via DI container
configManager := app.MustMake("config").(config.Manager)
```

### Dependency Injection in Services
```go
type MyService struct {
    config config.Manager
}

func NewMyService(app app.Application) *MyService {
    return &MyService{
        config: app.MustMake("config").(config.Manager),
    }
}
```

### Service Provider Configuration
```go
func (p *DatabaseServiceProvider) Register(container di.Container) error {
    return container.Singleton("database", func(container di.Container) (interface{}, error) {
        configManager := container.MustMake("config").(config.Manager)
        
        var dbConfig DatabaseConfig
        configManager.UnmarshalKey("database", &dbConfig)
        
        return newDatabase(dbConfig), nil
    })
}
```

## Configuration Lifecycle

### 1. Application Bootstrap
During application startup, the config service provider:
- Registers the config manager singleton
- Sets up default configuration paths
- Configures environment variable integration

### 2. Configuration Loading
Applications typically load configuration during bootstrap:
- Set configuration file name and paths
- Configure environment prefix
- Read configuration files
- Validate required settings

### 3. Runtime Access
Throughout the application lifecycle:
- Services access config via DI container
- Values are retrieved with type safety
- Configuration changes trigger watchers (if enabled)

### 4. Testing
During testing:
- Mock the config.Manager interface
- Set up test-specific configuration
- Isolate configuration state between tests

## Best Practices

### 1. Use Structured Configuration
Organize related configuration into nested structures:

```yaml
database:
  host: "localhost"
  port: 5432
  credentials:
    username: "user"
    password: "pass"
```

### 2. Provide Sensible Defaults
Set defaults for non-critical configuration:

```go
configManager.SetDefault("server.timeout", "30s")
configManager.SetDefault("cache.enabled", true)
```

### 3. Validate Critical Configuration
Validate required configuration at startup:

```go
if _, exists := configManager.GetString("database.host"); !exists {
    log.Fatal("Database host is required")
}
```

### 4. Use Environment Variables for Secrets
Override sensitive configuration with environment variables:

```bash
export FORK_DATABASE_PASSWORD=production-secret
export FORK_API_KEY=secret-api-key
```

### 5. Leverage Struct Unmarshaling
Use structs for complex configuration sections:

```go
type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
}

var dbConfig DatabaseConfig
configManager.UnmarshalKey("database", &dbConfig)
```

## Testing Strategy

### Unit Testing
Mock the config.Manager interface for unit tests:

```go
mockConfig := mocks.NewMockManager(t)
mockConfig.On("GetString", "key").Return("value", true)
```

### Integration Testing
Use real configuration with test-specific values:

```go
configManager.Set("database.host", "test-db")
configManager.Set("app.debug", true)
```

### Configuration Validation Testing
Test configuration validation logic:

```go
func TestConfigValidation(t *testing.T) {
    tests := []struct {
        name   string
        config map[string]interface{}
        valid  bool
    }{
        {"valid config", map[string]interface{}{"host": "localhost"}, true},
        {"missing host", map[string]interface{}{}, false},
    }
    // ...test implementation
}
```

## Error Handling

The config package uses Go's standard error handling patterns:

### File Errors
```go
if err := configManager.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // Handle missing file gracefully
        log.Println("Config file not found, using defaults")
    } else {
        // Handle other errors
        log.Fatal("Config error:", err)
    }
}
```

### Unmarshaling Errors
```go
var config MyConfig
if err := configManager.UnmarshalKey("section", &config); err != nil {
    return fmt.Errorf("invalid configuration: %w", err)
}
```

### Validation Errors
```go
func validateConfig(cfg config.Manager) error {
    if host, exists := cfg.GetString("database.host"); !exists || host == "" {
        return errors.New("database.host is required")
    }
    return nil
}
```
