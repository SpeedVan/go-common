package config

import (
	"strconv"

	"github.com/alpha-supsys/go-common/config"
	"github.com/alpha-supsys/go-common/type/collection/omap"
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

// GetInt todo
func (s *Config) GetInt(key string, _default int) int {
	if r, ok := s.cfg[s.Prefix+key]; ok {
		if ir, err := strconv.Atoi(r); err == nil {
			return ir
		}
	}
	return _default
}

// GetBool todo
func (s *Config) GetBool(key string, _default bool) bool {
	if r, ok := s.cfg[s.Prefix+key]; ok {
		if br, err := strconv.ParseBool(r); err == nil {
			return br
		}
	}
	return _default
}

// GetString todo
func (s *Config) GetString(key string, _default string) string {
	if result, ok := s.cfg[s.Prefix+key]; ok {
		return result
	}
	return _default
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
