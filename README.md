![coverage_badge.svg](docs/images/coverage_badge.svg)
# slog – Structured Logging for Go

A clean, idiomatic logger with syslog-style levels and a special `Crazy` dev-only level.

## Features

- Syslog-style log levels: `emerg`, `alert`, `crit`, `error`, `warn`, `notice`, `info`, `debug`
- Dev-only `Crazy` log level, only included with `-tags dev`
- Log level filtering: set the minimum level to log
- Simple API: use `slog.Info(...)`, `slog.Error(...)`, etc.

## Installation

```bash
go get github.com/TenforwardAB/slog
```

## Usage

```go
import "github.com/TenforwardAB/slog"

func main() {
	slog.SetLevel("info") // or "debug", "warn", etc.

	slog.Info("Application started")
	slog.Warn("This is a warning")
	slog.Crazy("This is a dev-only log") // Only works with -tags dev
}
```

## Supported log levels

| Name        | Alias(es)       | Severity (low → high) | Description                                           |
|-------------|------------------|------------------------|-------------------------------------------------------|
| `debug`     | –                | 0                      | Detailed internal debugging messages                 |
| `info`      | –                | 1                      | General system information                           |
| `notice`    | –                | 2                      | Significant but expected events                      |
| `warn`      | `warning`        | 3                      | Potential problems                                   |
| `error`     | –                | 4                      | Errors that require attention                        |
| `crit`      | `critical`       | 5                      | Critical conditions that may crash the system        |
| `alert`     | –                | 6                      | Immediate action needed                              |
| `emerg`     | `emergency`      | 7                      | System is unusable                                   |
| `crazy`     | *dev only*       | 0                      | Experimental/dev-only logs (only with `-tags dev`)   |

## Invalid log levels

If you call:

```go
slog.SetLevel("banana")
```

You’ll get:

```
invalid log level: banana
```

You can handle this:

```go
if err := slog.SetLevel(cfg.LogLevel); err != nil {
	log.Printf("invalid log level %q, falling back to info", cfg.LogLevel)
	slog.SetLevel(slog.LevelInfo)
}
```

## Dev-only logs

The `Crazy()` method is only compiled in if you build with:

```bash
go run -tags dev examples/main.go
```

In production, the `Crazy()` method becomes a no-op.

## Example

See [`examples/main.go`](examples/main.go) for a complete example.