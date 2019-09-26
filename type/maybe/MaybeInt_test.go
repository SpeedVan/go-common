package maybe

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func testType(flag bool) MaybeInt {
	if flag {
		return Integer(1)
	}

	return NullError
}

func Test(t *testing.T) {
	val1 := testType(false)
	val2 := testType(true)
	fmt.Println(reflect.TypeOf(val1))
	fmt.Println(reflect.TypeOf(val2))
	if v, ok := val1.(error); ok {
		fmt.Println(v)
	}
	if v, ok := val2.(Integer); ok {
		fmt.Println(v * 3)
	}

	switch t := reflect.TypeOf(errors.New("123")); t {
	case subtypeof(t, reflect.TypeOf(*NullError)):
		println("subtype:nullError")
	case subtypeof(t, reflect.TypeOf(Integer(1))):
		println("subtype:Integer")
	default:
		println("fail")
	}
}

func subtypeof(s reflect.Type, t reflect.Type) reflect.Type {
	println("s:" + s.Name() + ",t:" + t.Name())
	if s.ConvertibleTo(t) {
		return s
	}
	return t
}

func instanceof(v interface{}, t reflect.Type) bool {
	return false
}
