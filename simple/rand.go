package simple

import (
	"crypto/rand"
	"fmt"
)

func RandString(n int) string {
	randBytes := make([]byte, n)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
