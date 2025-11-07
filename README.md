# iotyper

A static analysis tool for Go that detects iota usage without explicit type specification.

## Overview

`iotyper` is a golangci-lint plugin that ensures `iota` constants have explicit types to improve code clarity and type safety.

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