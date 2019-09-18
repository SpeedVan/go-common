package maybe

import (
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
}
