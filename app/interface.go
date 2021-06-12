package app

import "github.com/alpha-supsys/go-common/log"

// App todo
type App interface {
	Run(log.Level) error
	SimpleRun() error
}
