package common

import (
	"fmt"
	"time"

	"github.com/SpeedVan/go-common/log"
)

type CommonLogger struct {
	log.Logger
	prefix string
	level  log.Level
}

func NewCommon(level log.Level) *CommonLogger {
	return &CommonLogger{
		prefix: "prefix",
		level:  level,
	}
}

func (s *CommonLogger) Log(level log.Level, msg string) {
	fmt.Println("[" + level.FormatString() + "]" + time.Now().Format("2006-01-02 15:04:05") + " : " + msg)
}

func (s *CommonLogger) Error(msg string) {
	if s.level <= log.Error {
		s.Log(log.Error, msg)
	}
}

func (s *CommonLogger) Info(msg string) {
	if s.level <= log.Info {
		s.Log(log.Info, msg)
	}
}

func (s *CommonLogger) Warn(msg string) {
	if s.level <= log.Warn {
		s.Log(log.Warn, msg)
	}
}

func (s *CommonLogger) Debug(msg string) {
	if s.level <= log.Debug {
		s.Log(log.Debug, msg)
	}
}

// LogF common log with format
func (s *CommonLogger) LogF(level log.Level, format string, arrs ...interface{}) {
	fmt.Printf("["+level.FormatString()+"]"+time.Now().Format("2006-01-02 15:04:05")+" : "+format+"\n", arrs...)
}

func (s *CommonLogger) ErrorF(format string, arrs ...interface{}) {
	if s.level <= log.Error {
		s.LogF(log.Error, format, arrs...)
	}
}

func (s *CommonLogger) InfoF(format string, arrs ...interface{}) {
	if s.level <= log.Info {
		s.LogF(log.Info, format, arrs...)
	}
}

func (s *CommonLogger) WarnF(format string, arrs ...interface{}) {
	if s.level <= log.Warn {
		s.LogF(log.Warn, format, arrs...)
	}
}

func (s *CommonLogger) DebugF(format string, arrs ...interface{}) {
	if s.level <= log.Debug {
		s.LogF(log.Debug, format, arrs...)
	}
}

func (s *CommonLogger) SetLevel(level log.Level) log.Logger {
	s.level = level
	return s
}
