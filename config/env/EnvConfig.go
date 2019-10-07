package env

import (
	"errors"
	"os"
	"strings"

	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/type/either"
)

type MaybeConfig either.Either

type EnvConfig struct {
	config.Config
	MaybeConfig
	Prefix  string
	configs map[string]string
}

type Error struct {
	MaybeConfig
	error
}

// WithPrefix todo
func (s *EnvConfig) WithPrefix(p string) config.Config {
	return &EnvConfig{
		Prefix:  s.Prefix + p,
		configs: s.configs,
	}
}

// LoadAll todo
func LoadAll() (config.Config, error) {
	envs := os.Environ()
	configs := map[string]string{}
	for _, item := range envs {
		pair := strings.SplitN(item, "=", 2)
		if len(pair) < 2 {
			return nil, errors.New("envConfig error:" + item)
		}
		configs[pair[0]] = pair[1]
	}
	return &EnvConfig{
		configs: configs,
	}, nil
}

// Get todo
func (s *EnvConfig) Get(name string) string {
	return s.configs[s.Prefix+name]
}

func (s *EnvConfig) GetInt(name string) int {
	return 0
}
