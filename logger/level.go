package logger

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) ShouldLog(loggerLevel Level) bool {
	return l >= loggerLevel
}
