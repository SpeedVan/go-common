package netcat

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	netcat := &Netcat{
		AddressAndPort: "0.0.0.0:2018",
		WhenAccept: func(s string) {
			fmt.Println(s)
		},
	}

	netcat.Run()
	defer netcat.Close()
}
