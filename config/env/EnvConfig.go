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
	configs map[string]string
}

type Error struct {
	MaybeConfig
	error
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

func (s *EnvConfig) Get(name string) string {
	return s.configs[name]
}

func (s *EnvConfig) GetInt(name string) int {
	return 0
}
