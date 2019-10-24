package config

import (
	"github.com/SpeedVan/go-common/type/collection/omap"
)

// Config todo
type Config interface {
	WithPrefix(string) Config
	Get(string) string
	// GetArr(string) []string
	GetMap(string) omap.Map
	ForEachArrayConfig(string, func(Config))
}
