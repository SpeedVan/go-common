package env

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/type/collection/omap"
	"github.com/SpeedVan/go-common/type/either"
)

var (
	reg          = regexp.MustCompile(`([a-zA-Z0-9][a-zA-Z0-9_]*?)_(\d+)`)
	LIST_HANDLER = "list"
)

type MaybeConfig either.Either

type EnvConfig struct {
	config.Config
	MaybeConfig
	Prefix       string
	OriginConfig *omap.CompositeMap // 这是基础的配置信息
	MapConfig    *omap.CompositeMap // 基础配置信息加工器 Tuple(result, handlerFunc)
}

func listHandler(k, v string, state *omap.CompositeMap) (*omap.CompositeMap, error) {
	err := state.RecursionSet(k, v)
	if err != nil {
		return nil, err
	}
	return state, nil
}

type Error struct {
	MaybeConfig
	error
}

// WithPrefix todo
func (s *EnvConfig) WithPrefix(p string) config.Config {
	return &EnvConfig{
		Prefix:       s.Prefix + p,
		OriginConfig: s.OriginConfig,
		MapConfig:    s.MapConfig,
	}
}

// LoadAll todo
func LoadAll() (config.Config, error) {
	return LoadAllWithoutPrefix("")
}

// LoadAllWithoutPrefix todo
func LoadAllWithoutPrefix(prefix string) (config.Config, error) {
	envs := os.Environ()
	config := &omap.CompositeMap{Data: make(omap.ComplexMap)}
	cacheConfig := &omap.CompositeMap{Data: make(omap.ComplexMap)}
	var err error
	for _, item := range envs {
		if strings.HasPrefix(item, prefix) {
			pair := strings.SplitN(strings.TrimPrefix(item, prefix), "=", 2)
			if len(pair) < 2 {
				fmt.Println(errors.New("envConfig error:" + item))
			} else {
				config.Set(pair[0], pair[1])
				cacheConfig, err = listHandler(pair[0], pair[1], cacheConfig)
				if err != nil {
					return nil, err
				}
				// config.Set(pair[0], pair[1])
			}
		}
	}
	return &EnvConfig{
		OriginConfig: config,
		MapConfig:    cacheConfig,
	}, nil
}

// func GetList(name string) []interface{} {
// 	for k, v := range s.configs {

// 	}
// }

// Get todo
func (s *EnvConfig) Get(name string) string {
	result, err := s.MapConfig.RecursionGet(s.Prefix + name)
	if err != nil {
		return ""
	}
	if str, ok := result.(string); ok {
		return str
	}
	return ""
}

func (s *EnvConfig) GetMap(name string) omap.Map {
	complexmap := s.MapConfig
	// fmt.Println(complexmap)
	v, err := complexmap.RecursionGet(s.Prefix + name)
	if err != nil {
		return nil
	}
	if result, ok := v.(*omap.CompositeMap); ok {
		return result
	}
	return nil
}

func (s *EnvConfig) GetCompositeMap(name string) *omap.CompositeMap {
	complexmap := s.MapConfig
	// fmt.Println(complexmap)
	v, err := complexmap.RecursionGet(name)
	if err != nil {
		return nil
	}
	if result, ok := v.(*omap.CompositeMap); ok {
		return result
	}
	return nil
}

func (s *EnvConfig) GetConfig(name string) config.Config {

	return &EnvConfig{
		Prefix: "",
	}
}

func (s *EnvConfig) ForEachArrayConfig(name string, handler func(config.Config)) {
	s.GetCompositeMap(s.Prefix + name).ForEach(func(k string, v interface{}) {
		if cmap, ok := v.(*omap.CompositeMap); ok {
			// fmt.Println(cmap)

			handler(&EnvConfig{
				Prefix: "",
				// OriginConfig: s.OriginConfig.F_map(func(k2 string, v2 interface{}) (string, interface{}) {
				// 	 return strings.TrimPrefix(s.Prefix+name+"_"+k+"_", k2), v
				// }).(*omap.CompositeMap), // 这是基础的配置信息
				MapConfig: cmap, // 基础配置信息加工器 Tuple(result, handlerFunc)
			})
		}
	})
}

func (s *EnvConfig) GetInt(name string) int {

	return 0
}

func (s *EnvConfig) String() string {
	return s.MapConfig.String()
}

func (s *EnvConfig) ToGolangStringMap() map[string]string {
	result := make(map[string]string)
	s.OriginConfig.ForEach(func(k string, v interface{}) {
		if strings.HasPrefix(k, s.Prefix) {
			result[strings.TrimPrefix(k, s.Prefix)] = fmt.Sprintf("%v", v)
		}
	})
	return result
}
