package env_test

import (
	"fmt"
	"testing"

	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/config/env"
)

func Test(t *testing.T) {
	conf, err := env.LoadAllWithoutPrefix("abc_")
	if err != nil {
		t.Error(err)
	}

	// c := conf.WithPrefix("CHECK_CONFIG_")

	conf.ForEachArrayConfig("CHECK_CONFIG", func(c config.Config) {
		fmt.Println(c.Get("name"))
	})

	conf.ForEachArrayConfig("CHECK_CONFIG", func(c config.Config) {
		fmt.Println(c.Get("name"))
	})

	fmt.Println(conf.WithPrefix("CHECK_").ToGolangStringMap())
}
