# AGENTS.md

## Project
Simple Go semantic logger with ANSI color output.

## Commands
- `go build ./...` - build all packages
- `go test ./...` - run tests
- `go vet ./...` - lint

## Structure
- `logger/` - Logger interface, implementations (ConsoleLogger, FileLogger), and Level type
- `ansi/` - ANSI color constants

## Notes
- Module: `lucasaladino.com/semanticlog`
- Go version: 1.25.5
- No tests currently exist
- Output goes to stderr (ConsoleLogger writes to os.Stderr by default)

## Logging Levels
- `logger.LevelDebug` - log everything (default)
- `logger.LevelInfo` - log info, warn, error
- `logger.LevelWarn` - log warn, error
- `logger.LevelError` - log only error
- Set in constructor: `NewConsoleLogger(logger.LevelInfo)`, `NewFileLogger("path", "label", logger.LevelWarn)`
- Or via setter: `logger.SetLevel(logger.LevelError)`
