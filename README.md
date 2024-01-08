# Safe
> Safe function execution without panic

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/once)](https://pkg.go.dev/github.com/go-zoox/once)
[![Build Status](https://github.com/go-zoox/once/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/once/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/once)](https://goreportcard.com/report/github.com/go-zoox/once)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/safe/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/safe?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/safe.svg)](https://github.com/go-zoox/once/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/safe.svg?label=Release)](https://github.com/go-zoox/once/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/once
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/once"
)

func main(t *testing.T) {
	loadConfig := func() {
		panic("load config panic")
	}

	if err := safe.Do(loadConfig); err != nil {
		log.Error(err)
	}
}
```

## Inspired By
* [kenkyu392/go-safe](https://github.com/kenkyu392/go-safe) - Provides a sandbox for the safe execution of panic-inducing programs
* [go-zoox/retry](https://github.com/andskur/argon2-hashing) - Catch Panic In Retries

## License
GoZoox is released under the [MIT License](./LICENSE).
