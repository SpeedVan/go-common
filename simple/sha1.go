package simple

import (
	"crypto/sha1"
	"fmt"
)

func Sha1(abs string) string {
	cry := sha1.New()
	cry.Write([]byte(abs))
	return fmt.Sprintf("%x", cry.Sum(nil))
}
