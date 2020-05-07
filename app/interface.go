package app

import "github.com/SpeedVan/go-common/log"

// App todo
type App interface {
	Run(log.Level) error
	SimpleRun() error
}
