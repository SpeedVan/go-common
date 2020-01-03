package net

import (
	"net"
)

// CheckPort todo
func CheckPort(port string) error {

	ln, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		return err
	}
	defer func() {
		ln.Close()
	}()

	return nil
}
