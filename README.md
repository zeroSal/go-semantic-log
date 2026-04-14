# go-semantic-log
The modular, injectable, meaningfully colorful logger for Go

## Install

```bash
go get github.com/zeroSal/go-semantic-log
```

## Usage

### Console Logger

```go
import (
    "github.com/zeroSal/go-semantic-log/logger"
)

logger := logger.NewConsoleLogger()
logger.Info("Hello world")

// With logging level
logger := logger.NewConsoleLogger(logger.LevelWarn)
logger.SetLevel(logger.LevelError)
```

### File Logger

```go
logger := logger.NewFileLogger("/var/log/app.log", "myapp")

// Open the file
if err := logger.Init(); err != nil {
    // handle error
}
defer logger.Close() // Closes the file

logger.Info("Hello world")

// With logging level
logger := logger.NewFileLogger("/var/log/app.log", "myapp", logger.LevelInfo)
logger.SetLevel(logger.LevelWarn)
```

## Logging Levels

- `logger.LevelDebug` - log everything (default)
- `logger.LevelInfo` - log info, warn, error
- `logger.LevelWarn` - log warn, error  
- `logger.LevelError` - log only error

Set via constructor or `SetLevel()`.

## Uber FX Integration

To use with FX lifecycle:

```go
func provideLogger(lc fx.Lifecycle) *logger.FileLogger {
    l := logger.NewFileLogger("/var/log/error.log", "error")
    
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            return l.Init()
        },
        OnStop: func(ctx context.Context) error {
            return l.Close()
        },
    })
    
    return l
}
```

Then inject it into your handlers:

```go
fx.Provide(provideLogger),
fx.Invoke(func(l *logger.FileLogger) { ... }),
```
