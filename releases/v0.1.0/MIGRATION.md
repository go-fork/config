# Migration Guide - v0.1.0

## Overview
This is the initial release of go-fork/config package. No migration is needed as this is the first version.

## Installation

### Step 1: Install Package
```bash
go get go.fork.vn/config@v0.1.0
```

### Step 2: Basic Usage
```go
package main

import (
    "go.fork.vn/config"
    "go.fork.vn/di"
)

func main() {
    // Create DI application
    app := di.New()
    
    // Register config service provider
    app.Register(config.NewServiceProvider())
    
    // Get configuration manager
    container := app.Container()
    cfg := container.MustMake("config").(config.Manager)
    
    // Use configuration
    setting := cfg.Get("app.name", "default-app")
    fmt.Println("App name:", setting)
}
```

## Getting Started
- Check the [documentation](https://pkg.go.dev/go.fork.vn/config@v0.1.0)
- Review examples in the repository
- Read the README.md for detailed usage instructions

---
**Welcome to go-fork/config!** This is the beginning of a powerful configuration management solution.
