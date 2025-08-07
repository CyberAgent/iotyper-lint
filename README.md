# iotyper

A static analysis tool for Go that detects iota usage without explicit type specification.

## Overview

`iotyper` is a golangci-lint plugin that ensures `iota` constants have explicit types to improve code clarity and type safety.

## Installation

### Step 1: Create Plugin Configuration

Create a `.custom-gcl.yml` file in your project root:

```yaml
# This has to be >= v1.57.0 for module plugin system support.
version: v1.57.0
plugins:
  - module: 'github.com/CyberAgent/iotyper-lint'
    import: 'github.com/CyberAgent/iotyper-lint/linters'
    version: latest # Or a fixed version for reproducible builds.
```

### Step 2: Configure golangci-lint

Add to your `.golangci.yml`:

```yaml
linters:
  enable:
    - iotyper
```

### Step 3: Build Custom Binary

Build a custom golangci-lint binary with the plugin:

```bash
golangci-lint custom
```

### Step 4: Run the Linter

```bash
./custom-gcl run ./...
```

## Usage

### What Gets Flagged

`iotyper` reports when `iota` is used in const declarations without an explicit type:

```go
// ❌ Bad: will be flagged
const (
    StatusPending = iota  // Error: iota used without type specification
    StatusActive
    StatusClosed
)
```

### How to Fix

Add an explicit type to your iota declaration:

```go
// ✅ Good: type explicitly specified
const (
    StatusPending int = iota
    StatusActive
    StatusClosed
)
```

### Suppressing False Positives

You can suppress the linter using `//nolint` comments:

```go
const (
    Value1 = iota  //nolint:iotyper
    Value2 = iota  //nolint:all
)
```

## License

See [LICENSE](LICENSE) file for details.