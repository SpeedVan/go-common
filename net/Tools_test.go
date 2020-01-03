package net

import (
	"fmt"
	"testing"
)

func Test(_ *testing.T) {
	err := CheckPort("8888")
	if err != nil {
		fmt.Println("error:" + err.Error())
	} else {
		fmt.Println("ok!")
	}
}
