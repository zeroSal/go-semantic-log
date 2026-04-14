package logger

import (
	"fmt"
	"os"
	"sync"

	"lucasaladino.com/semanticlog/ansi"
)

var _ LoggerInterface = (*FileLogger)(nil)

type FileLogger struct {
	path  string
	label string
	level Level

	file *os.File
	mu   sync.Mutex
}

func NewFileLogger(
	path,
	label string,
	level ...Level,
) *FileLogger {
	l := &FileLogger{
		path:  path,
		label: label,
	}
	if len(level) > 0 {
		l.level = level[0]
	}
	return l
}

func (l *FileLogger) SetLevel(level Level) {
	l.level = level
}

func (l *FileLogger) GetLevel() Level {
	return l.level
}

func (l *FileLogger) Init() error {
	f, err := os.OpenFile(
		l.path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		return err
	}

	l.file = f

	return nil
}

func (l *FileLogger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}

	return nil
}

func (l *FileLogger) log(prefix, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		l.logError("Logger "+l.label+" not initialized. Call Init() first.", nil)
		return
	}

	_, err := fmt.Fprintf(l.file, "%s %s\n", prefix, msg)
	if err != nil {
		l.logError("Write failed", err)
		fmt.Fprintf(os.Stderr, "[ERROR] write failed: %v\n", err)
	}
}

func (l *FileLogger) GetIdentifier() string {
	return l.label
}

func (l *FileLogger) logError(msg string, err error) {
	errorString := ""
	if err != nil {
		errorString = err.Error()
	}

	_, _ = fmt.Fprintf(
		os.Stderr,
		"%s%s %s%s%s\n",
		ansi.Red,
		"[ERROR]",
		msg,
		errorString,
		ansi.Reset,
	)
}

func (l *FileLogger) Debug(msg string) {
	if !LevelDebug.ShouldLog(l.level) {
		return
	}
	l.log("[DEBUG]", msg)
}

func (l *FileLogger) Info(msg string) {
	if !LevelInfo.ShouldLog(l.level) {
		return
	}
	l.log("[INFO]", msg)
}

func (l *FileLogger) Warn(msg string) {
	if !LevelWarn.ShouldLog(l.level) {
		return
	}
	l.log("[WARNING]", msg)
}

func (l *FileLogger) Error(msg string) {
	if !LevelError.ShouldLog(l.level) {
		return
	}
	l.log("[ERROR]", msg)
}

func (l *FileLogger) List(msgs []string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		l.logError("Logger "+l.label+" not initialized. Call Init() first.", nil)
		return
	}

	for _, msg := range msgs {
		_, err := fmt.Fprintf(l.file, " · %s\n", msg)
		if err != nil {
			l.logError("Write failed", err)
			fmt.Fprintf(os.Stderr, "[ERROR] write failed: %v\n", err)
		}
	}
}
