package logger

import (
	"io"
)

type LoggerInterface interface {
	io.Closer
	Init() error
	SetLevel(level Level)
	GetLevel() Level
	GetIdentifier() string
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	List([]string)
}
