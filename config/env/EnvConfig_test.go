package env_test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/alpha-ss/go-common/config"
	"github.com/alpha-ss/go-common/config/env"
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

type TestStruct struct {
}

func (s *TestStruct) TestA(a, b string) {

}

func (s *TestStruct) fm(a, b string) {

}

type AllMethodHandler func(string, string)

type TestF map[string]interface{}

func Test2(t *testing.T) {
	s := &TestStruct{}

	m := make(TestF)
	m["TestA"] = runtime.FuncForPC(reflect.ValueOf(AllMethodHandler(s.TestA)).Pointer()).Name()
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(AllMethodHandler(s.fm)).Pointer()).Name())
}
