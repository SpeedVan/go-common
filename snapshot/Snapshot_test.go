package snapshot

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/alpha-supsys/go-common/simple"
)

func TestSnapshot(t *testing.T) {
	init := &sync.Map{}
	ss := NewSnapshot(init, func(e *Event, data interface{}, s string) (interface{}, error) {
		mapdata := data.(*sync.Map)
		mapdata.Store(e.data.String(), s)

		return data, nil
	})

	ss.Run()

	go func() {
		for {
			ss.PutStr(simple.RandString(16))
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("data:")
			(ss.data.(*sync.Map)).Range(func(key, value interface{}) bool {
				fmt.Println(key, value)
				return true
			})
			time.Sleep(3 * time.Second)
		}
	}()

	fmt.Scanln()
}
