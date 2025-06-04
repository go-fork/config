# Config Package Documentation

The config package provides robust configuration management for Fork Framework applications, integrating seamlessly with the dependency injection container and service provider pattern.

## Quick Navigation

- **[Overview](overview.md)** - Architecture and concepts
- **[Usage Guide](usage.md)** - Practical examples and patterns
- **[API Reference](reference.md)** - Complete interface documentation

## Quick Start

```go
package main

import (
    "log"
    "go.fork.vn/app"
)

func main() {
    // Create Fork application
    app := app.NewApplication()
    
    // Access config through DI container
    configManager := app.MustMake("config").(config.Manager)
    
    // Configure and read
    configManager.SetConfigName("app")
    configManager.AddConfigPath("./configs")
    configManager.SetConfigType("yaml")
    
    if err := configManager.ReadInConfig(); err != nil {
        log.Printf("Config warning: %v", err)
    }
    
    // Use configuration
    if appName, exists := configManager.GetString("app.name"); exists {
        log.Printf("Application: %s", appName)
    }
    
    app.Run()
}
```

## Key Features

- **Framework Integration**: Seamless integration with Fork Framework's DI container
- **Multiple Formats**: Support for YAML, JSON, TOML, and environment variables
- **Type Safety**: Type-safe value retrieval with existence checks
- **Live Reload**: Configuration watching with automatic reloading
- **Environment Override**: Automatic environment variable precedence
- **Validation**: Built-in validation patterns and error handling

## Installation

```bash
go get go.fork.vn/config@latest
```

The config service is automatically registered when using Fork Framework.

## Package Structure

```
config/
├── docs/
│   ├── index.md      # This file
│   ├── overview.md   # Architecture and concepts
│   ├── usage.md      # Usage examples
│   └── reference.md  # API documentation
├── mocks/            # Generated mocks
├── manager.go        # Core implementation
├── service_provider.go # Framework integration
└── doc.go           # Package documentation
```

## Getting Help

- Read the [Usage Guide](usage.md) for practical examples
- Check the [API Reference](reference.md) for detailed interface documentation
- Review the [Overview](overview.md) for architectural concepts
