package mock

import (
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/type/collection/omap"
)

// Config todo
type Config struct {
	Prefix string
	cfg    map[string]string
	config.Config
}

// WithPrefix todo
func (s *Config) WithPrefix(p string) config.Config {
	return &Config{
		Prefix: s.Prefix + p,
		cfg:    s.cfg,
	}
}

// Get todo
func (s *Config) Get(key string) string {
	if r, ok := s.cfg[s.Prefix+key]; ok {
		return r
	}
	return ""
}

// GetMap todo
func (s *Config) GetMap(string) omap.Map {

	return nil
}

// ForEachArrayConfig todo
func (s *Config) ForEachArrayConfig(string, func(config.Config)) {

	return
}

// ToGolangStringMap todo
func (s *Config) ToGolangStringMap() map[string]string {

	return map[string]string{}
}

// New todo
func New(cfg map[string]string) config.Config {
	return &Config{
		Prefix: "",
		cfg:    cfg,
	}
}
