package common

import (
	"fmt"
	"time"
	"github.com/SpeedVan/go-common/log"
)

type CommonLogger struct {
	log.Logger
	level log.Level
}

func NewCommon(level log.Level) *CommonLogger {
	return &CommonLogger{
		level: level,
	}
}

func (s *CommonLogger) Log(level log.Level, format string, arrs ...interface{}) {
	fmt.Printf("["+level.FormatString()+"]"+time.Now().Format("2006-01-02 15:04:05")+" : " + format + "/n", arrs...)
}

func (s *CommonLogger) Error(format string, arrs ...interface{}) {
	if s.level > log.Error {
		s.Log(log.Error, format, arrs)
	}
}

func (s *CommonLogger) Info(format string, arrs ...interface{}) {
	if s.level > log.Info {
		s.Log(log.Info, format, arrs)
	}
}

func (s *CommonLogger) Warn(format string, arrs ...interface{}) {
	if s.level > log.Warn {
		s.Log(log.Warn, format, arrs)
	}
}

func (s *CommonLogger) Debug(format string, arrs ...interface{}) {
	if s.level > log.Debug {
		s.Log(log.Debug, format, arrs)
	}
}