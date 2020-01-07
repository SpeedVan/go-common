package log

import (
	"fmt"
	"time"
)

type CommonLogger struct {
	Logger
	prefix string
	level  Level
}

func NewCommon(level Level) *CommonLogger {
	return &CommonLogger{
		prefix: "prefix",
		level:  level,
	}
}

func (s *CommonLogger) Log(level Level, msg string) {
	fmt.Println("[" + level.FormatString() + "]" + time.Now().Format("2006-01-02 15:04:05") + " : " + msg)
}

func (s *CommonLogger) Error(msg string) {
	if s.level <= Error {
		s.Log(Error, msg)
	}
}

func (s *CommonLogger) Info(msg string) {
	if s.level <= Info {
		s.Log(Info, msg)
	}
}

func (s *CommonLogger) Warn(msg string) {
	if s.level <= Warn {
		s.Log(Warn, msg)
	}
}

func (s *CommonLogger) Debug(msg string) {
	if s.level <= Debug {
		s.Log(Debug, msg)
	}
}

// LogF common log with format
func (s *CommonLogger) LogF(level Level, format string, arrs ...interface{}) {
	fmt.Printf("["+level.FormatString()+"]"+time.Now().Format("2006-01-02 15:04:05")+" : "+format+"\n", arrs...)
}

func (s *CommonLogger) ErrorF(format string, arrs ...interface{}) {
	if s.level <= Error {
		s.LogF(Error, format, arrs...)
	}
}

func (s *CommonLogger) InfoF(format string, arrs ...interface{}) {
	if s.level <= Info {
		s.LogF(Info, format, arrs...)
	}
}

func (s *CommonLogger) WarnF(format string, arrs ...interface{}) {
	if s.level <= Warn {
		s.LogF(Warn, format, arrs...)
	}
}

func (s *CommonLogger) DebugF(format string, arrs ...interface{}) {
	if s.level <= Debug {
		s.LogF(Debug, format, arrs...)
	}
}

func (s *CommonLogger) SetLevel(level Level) Logger {
	s.level = level
	return s
}
