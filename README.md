# slogic

🧠 _Composable filtering logic for [`log/slog`](https://pkg.go.dev/log/slog)_

[![Test](https://github.com/lukephillippi/slogic/actions/workflows/test.yaml/badge.svg)](https://github.com/lukephillippi/slogic/actions/workflows/test.yaml)
[![Go Reference](https://pkg.go.dev/badge/go.luke.ph/slogic.svg)](https://pkg.go.dev/go.luke.ph/slogic)

## Overview

The `slogic` package provides surgical control over what gets logged in your Go applications.

Built on top of [the standard library's `log/slog` package](https://pkg.go.dev/log/slog), it provides a composable filtering system that lets you:

- ✅ Dynamically filter out logs based on level, time, message content, and/or any key-value attribute
- ✅ Formulate bespoke filtering rules with logical operators (`And`, `Or`, `Not`)
- ✅ Apply filters to any `log/slog` `Handler` implementation
- ✅ Implement custom filters via a simple `Filter` interface

It's lightweight, dependency-free, and integrates seamlessly with any existing `log/slog`-based logging setup.

## Installing

1. First, use `go get` to install the latest version of the package:

   ```shell
   go get -u go.luke.ph/slogic@latest
   ```

1. Next, include the package in your application:

   ```go
   import "go.luke.ph/slogic"
   ```

## Usage

The `slogic` package provides the core filtering functionality, while the `slogic/filter` package offers a rich set of pre-built `Filter`s suitable for a variety of common use cases.

Combine these with logical operators (`And`, `Or`, `Not`) to create sophisticated filtering rules:

```go
// 1. Keep all ERROR logs
// 2. Keep latency WARN logs
// 3. Filter all other ≤ WARN logs
handler := slogic.NewHandler(
    slog.NewTextHandler(os.Stdout, nil),
    slogic.And(
        slogic.Not(
            slogic.Or(
                filter.IfLevelEquals(slog.LevelError),
                slogic.And(
                    filter.IfLevelEquals(slog.LevelWarn),
                    filter.IfAttrExists("latency_ms"),
                ),
            ),
        ),
        filter.IfLevelAtMost(slog.LevelWarn),
    ),
)

slog.SetDefault(slog.New(handler))
```

## License

The package is released under [the Unlicense license](./LICENSE.md).

## References

- [pkg.go.dev/log/slog](https://pkg.go.dev/log/slog)
- [Structured Logging with slog](https://go.dev/blog/slog)
- [A Guide to Writing `slog` Handlers](https://github.com/golang/example/blob/master/slog-handler-guide/README.md)
