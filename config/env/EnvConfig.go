package env

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alpha-ss/go-common/config"
	"github.com/alpha-ss/go-common/type/collection/omap"
	"github.com/alpha-ss/go-common/type/either"
)

var (
	reg = regexp.MustCompile(`([a-zA-Z0-9][a-zA-Z0-9_]*?)_(\d+)`)
	// LIST_HANDLER todo
	LIST_HANDLER = "list"
)

// MaybeConfig todo
type MaybeConfig either.Either

// EnvConfig todo
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

// Error todo
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

// Get todo
func (s *EnvConfig) Get(name string) string {
	str, err := s.MapConfig.RecursionGet(s.Prefix + name)
	if err != nil {
		return ""
	}
	if result, ok := str.(string); ok {
		return result
	}
	return ""
}

// GetString todo
func (s *EnvConfig) GetString(name string, _default string) string {
	str, err := s.MapConfig.RecursionGet(s.Prefix + name)
	if err != nil {
		return _default
	}
	if result, ok := str.(string); ok {
		return result
	}
	return _default
}

// GetInt todo
func (s *EnvConfig) GetInt(name string, _default int) int {
	str, err := s.MapConfig.RecursionGet(s.Prefix + name)
	if err != nil {
		return _default
	}
	intResult, err := strconv.Atoi(fmt.Sprint(str))
	if err != nil {
		return _default
	}
	return intResult
}

// GetBool todo
func (s *EnvConfig) GetBool(name string, _default bool) bool {
	str, err := s.MapConfig.RecursionGet(s.Prefix + name)
	if err != nil {
		return _default
	}
	boolResult, err := strconv.ParseBool(fmt.Sprint(str))
	if err != nil {
		return _default
	}
	return boolResult
}

// GetMap todo
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

// GetCompositeMap todo
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

// GetConfig todo
func (s *EnvConfig) GetConfig(name string) config.Config {

	return &EnvConfig{
		Prefix: "",
	}
}

// ForEachArrayConfig todo
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

func (s *EnvConfig) String() string {
	return s.MapConfig.String()
}

// ToGolangStringMap todo
func (s *EnvConfig) ToGolangStringMap() map[string]string {
	result := make(map[string]string)
	s.OriginConfig.ForEach(func(k string, v interface{}) {
		if strings.HasPrefix(k, s.Prefix) {
			result[strings.TrimPrefix(k, s.Prefix)] = fmt.Sprintf("%v", v)
		}
	})
	return result
}
