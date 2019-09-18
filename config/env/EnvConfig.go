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

func LoadAll() MaybeConfig {
	envs := os.Environ()
	configs := map[string]string{}
	for _, item := range envs {
		pair := strings.SplitN(item, "=", 1)
		if len(pair) < 2 {
			return errors.New("")
		}
		configs[pair[0]] = pair[1]
	}
	return &EnvConfig{
		configs: configs,
	}
}

func (s *EnvConfig) Get(name string) string {
	return s.configs[name]
}

func (s *EnvConfig) GetInt(name string) int {

}
