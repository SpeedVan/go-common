package config

import (
	"github.com/alpha-supsys/go-common/type/collection/omap"
)

// Config todo
type Config interface {
	WithPrefix(string) Config
	Get(string) string
	GetInt(string, int) int
	GetBool(string, bool) bool
	GetString(string, string) string
	// GetArr(string) []string
	GetMap(string) omap.Map
	ForEachArrayConfig(string, func(Config))
	ToGolangStringMap() map[string]string
}
