# Config Package API Reference

## Interfaces

### Manager

The `Manager` interface provides all configuration management functionality.

```go
type Manager interface {
    // Configuration file operations
    SetConfigFile(in string)
    SetConfigName(in string)  
    SetConfigType(in string)
    AddConfigPath(in string)
    
    // Reading configuration
    ReadInConfig() error
    MergeInConfig() error
    MergeConfig(io.Reader) error
    
    // Value retrieval
    GetString(key string) (string, bool)
    GetInt(key string) (int, bool)
    GetInt32(key string) (int32, bool)
    GetInt64(key string) (int64, bool)
    GetUint(key string) (uint, bool)
    GetUint32(key string) (uint32, bool)
    GetUint64(key string) (uint64, bool)
    GetFloat64(key string) (float64, bool)
    GetBool(key string) (bool, bool)
    GetTime(key string) (time.Time, bool)
    GetDuration(key string) (time.Duration, bool)
    
    // Slice operations
    GetSlice(key string) ([]interface{}, bool)
    GetStringSlice(key string) ([]string, bool)
    GetIntSlice(key string) ([]int, bool)
    
    // Map operations
    GetMap(key string) (map[string]interface{}, bool)
    GetStringMap(key string) (map[string]interface{}, bool)
    GetStringMapString(key string) (map[string]string, bool)
    GetStringMapStringSlice(key string) (map[string][]string, bool)
    
    // Unmarshaling
    Unmarshal(rawVal interface{}) error
    UnmarshalKey(key string, rawVal interface{}) error
    
    // Configuration management
    Set(key string, value interface{})
    SetDefault(key string, value interface{})
    Has(key string) bool
    
    // Environment variables
    AutomaticEnv()
    SetEnvPrefix(in string)
    BindEnv(input ...string) error
    
    // Configuration watching
    WatchConfig()
    OnConfigChange(run func(in fsnotify.Event))
    
    // Writing configuration
    WriteConfig() error
    WriteConfigAs(filename string) error
    SafeWriteConfigAs(filename string) error
}
```

## Methods Documentation

### Configuration File Operations

#### SetConfigFile(in string)
Sets the exact path to the configuration file.

**Parameters:**
- `in` - Full path to configuration file

**Example:**
```go
configManager.SetConfigFile("/etc/myapp/config.yaml")
```

#### SetConfigName(in string)
Sets the name of the configuration file (without extension).

**Parameters:**
- `in` - Configuration file name without extension

**Example:**
```go
configManager.SetConfigName("app")  // Looks for app.yaml, app.json, etc.
```

#### SetConfigType(in string)
Sets the type of configuration file.

**Parameters:**
- `in` - File type ("yaml", "json", "toml", etc.)

**Example:**
```go
configManager.SetConfigType("yaml")
```

#### AddConfigPath(in string)
Adds a path to search for configuration files.

**Parameters:**
- `in` - Directory path to search

**Example:**
```go
configManager.AddConfigPath("./configs")
configManager.AddConfigPath("/etc/myapp")
```

### Reading Configuration

#### ReadInConfig() error
Reads configuration from the configured file.

**Returns:**
- `error` - Error if file cannot be read or parsed

**Example:**
```go
if err := configManager.ReadInConfig(); err != nil {
    log.Fatal("Cannot read config:", err)
}
```

#### MergeInConfig() error
Merges configuration from file with existing configuration.

**Returns:**
- `error` - Error if merge fails

#### MergeConfig(io.Reader) error
Merges configuration from a reader with existing configuration.

**Parameters:**
- `io.Reader` - Reader containing configuration data

**Returns:**
- `error` - Error if merge fails

### Value Retrieval Methods

All value retrieval methods follow the same pattern: they return the value and a boolean indicating whether the key exists.

#### GetString(key string) (string, bool)
Retrieves a string value.

**Parameters:**
- `key` - Configuration key using dot notation

**Returns:**
- `string` - The string value
- `bool` - True if key exists, false otherwise

**Example:**
```go
if host, exists := configManager.GetString("database.host"); exists {
    fmt.Printf("Database host: %s\n", host)
}
```

#### GetInt(key string) (int, bool)
Retrieves an integer value.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `int` - The integer value
- `bool` - True if key exists and can be converted to int

**Example:**
```go
if port, exists := configManager.GetInt("server.port"); exists {
    server.Listen(port)
}
```

#### GetBool(key string) (bool, bool)
Retrieves a boolean value.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `bool` - The boolean value
- `bool` - True if key exists and can be converted to bool

**Example:**
```go
if debug, exists := configManager.GetBool("app.debug"); exists && debug {
    enableDebugMode()
}
```

#### GetDuration(key string) (time.Duration, bool)
Retrieves a duration value. Accepts strings like "30s", "5m", "1h".

**Parameters:**
- `key` - Configuration key

**Returns:**
- `time.Duration` - The duration value
- `bool` - True if key exists and can be parsed as duration

**Example:**
```go
if timeout, exists := configManager.GetDuration("server.timeout"); exists {
    server.SetTimeout(timeout)
}
```

#### GetTime(key string) (time.Time, bool)
Retrieves a time value.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `time.Time` - The time value
- `bool` - True if key exists and can be parsed as time

### Slice Operations

#### GetStringSlice(key string) ([]string, bool)
Retrieves a slice of strings.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `[]string` - Slice of string values
- `bool` - True if key exists

**Example:**
```go
if hosts, exists := configManager.GetStringSlice("database.replicas"); exists {
    for _, host := range hosts {
        pool.AddReplica(host)
    }
}
```

#### GetIntSlice(key string) ([]int, bool)
Retrieves a slice of integers.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `[]int` - Slice of integer values
- `bool` - True if key exists

### Map Operations

#### GetStringMap(key string) (map[string]interface{}, bool)
Retrieves a map with string keys and interface{} values.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `map[string]interface{}` - Map of values
- `bool` - True if key exists

#### GetStringMapString(key string) (map[string]string, bool)
Retrieves a map with string keys and string values.

**Parameters:**
- `key` - Configuration key

**Returns:**
- `map[string]string` - Map of string values
- `bool` - True if key exists

**Example:**
```go
if headers, exists := configManager.GetStringMapString("http.headers"); exists {
    for name, value := range headers {
        req.Header.Set(name, value)
    }
}
```

### Unmarshaling

#### Unmarshal(rawVal interface{}) error
Unmarshals the entire configuration into a struct.

**Parameters:**
- `rawVal` - Pointer to struct to unmarshal into

**Returns:**
- `error` - Error if unmarshaling fails

**Example:**
```go
type Config struct {
    App      AppConfig      `mapstructure:"app"`
    Database DatabaseConfig `mapstructure:"database"`
}

var config Config
if err := configManager.Unmarshal(&config); err != nil {
    log.Fatal("Cannot unmarshal config:", err)
}
```

#### UnmarshalKey(key string, rawVal interface{}) error
Unmarshals a specific configuration section into a struct.

**Parameters:**
- `key` - Configuration key for the section
- `rawVal` - Pointer to struct to unmarshal into

**Returns:**
- `error` - Error if unmarshaling fails

**Example:**
```go
type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Username string `mapstructure:"username"`
}

var dbConfig DatabaseConfig
if err := configManager.UnmarshalKey("database", &dbConfig); err != nil {
    log.Fatal("Cannot unmarshal database config:", err)
}
```

### Configuration Management

#### Set(key string, value interface{})
Sets a configuration value programmatically.

**Parameters:**
- `key` - Configuration key
- `value` - Value to set

**Example:**
```go
configManager.Set("app.version", "2.0.0")
configManager.Set("features.beta", true)
```

#### SetDefault(key string, value interface{})
Sets a default value for a configuration key.

**Parameters:**
- `key` - Configuration key
- `value` - Default value

**Example:**
```go
configManager.SetDefault("server.port", 8080)
configManager.SetDefault("app.debug", false)
```

#### Has(key string) bool
Checks if a configuration key exists.

**Parameters:**
- `key` - Configuration key to check

**Returns:**
- `bool` - True if key exists

**Example:**
```go
if configManager.Has("database.host") {
    // Database configuration is present
}
```

### Environment Variables

#### AutomaticEnv()
Enables automatic environment variable lookup.

**Example:**
```go
configManager.AutomaticEnv()
// Now FORK_DATABASE_HOST overrides database.host
```

#### SetEnvPrefix(in string)
Sets the prefix for environment variables.

**Parameters:**
- `in` - Environment variable prefix

**Example:**
```go
configManager.SetEnvPrefix("MYAPP")
// Now MYAPP_DATABASE_HOST overrides database.host
```

#### BindEnv(input ...string) error
Explicitly binds environment variables to configuration keys.

**Parameters:**
- `input` - Key and optional environment variable name

**Returns:**
- `error` - Error if binding fails

**Example:**
```go
configManager.BindEnv("database.password", "DB_PASSWORD")
configManager.BindEnv("debug") // Uses DEBUG env var
```

### Configuration Watching

#### WatchConfig()
Starts watching the configuration file for changes.

**Example:**
```go
configManager.WatchConfig()
```

#### OnConfigChange(run func(in fsnotify.Event))
Sets a callback function to run when configuration changes.

**Parameters:**
- `run` - Function to execute on configuration change

**Example:**
```go
configManager.OnConfigChange(func(e fsnotify.Event) {
    log.Printf("Config changed: %s", e.Name)
    reloadServices()
})
```

### Writing Configuration

#### WriteConfig() error
Writes current configuration to the configured file.

**Returns:**
- `error` - Error if write fails

#### WriteConfigAs(filename string) error
Writes current configuration to a specific file.

**Parameters:**
- `filename` - Target file path

**Returns:**
- `error` - Error if write fails

#### SafeWriteConfigAs(filename string) error
Writes configuration to file only if it doesn't exist.

**Parameters:**
- `filename` - Target file path

**Returns:**
- `error` - Error if write fails or file exists

## Constants and Types

### Supported Configuration Formats
- `yaml` / `yml` - YAML format
- `json` - JSON format
- `toml` - TOML format
- `ini` - INI format
- `properties` - Java properties format

### Error Types
The package may return various error types:
- `viper.ConfigFileNotFoundError` - Configuration file not found
- `viper.ConfigParseError` - Configuration parsing error
- Standard Go errors for other conditions

## Usage with Fork Framework

### Service Provider Integration
```go
// The config service is automatically available
app := app.NewApplication()
configManager := app.MustMake("config").(config.Manager)
```

### Testing
```go
import "go.fork.vn/config/mocks"

func TestMyService(t *testing.T) {
    mockConfig := mocks.NewMockManager(t)
    mockConfig.On("GetString", "database.host").Return("localhost", true)
    
    // Use mockConfig in tests
}
```
