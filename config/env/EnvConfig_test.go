package env_test

import (
	"fmt"
	"testing"

	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/config/env"
)

func Test(t *testing.T) {
	conf, err := env.LoadAllWithPrefix("abc_")
	if err != nil {
		t.Error(err)
	}

	conf.ForEachArrayConfig("CHECK_CONFIG", func(c config.Config) {
		fmt.Println(c.Get("name"))
	})

}
