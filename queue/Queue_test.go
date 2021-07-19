package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue(3)

	routine1 := func() {
		for i := 0; i < 10000; i = i + 2 {
			q.Put(i)
		}
	}

	routine2 := func() {
		for i := 1; i < 10000; i = i + 2 {
			q.Put(i)
		}
	}

	t1 := time.Now()
	go routine1()
	go routine2()

	for i := 0; i < 10000; i++ {
		// q.Get() // 若测试中，单独使用这行，将不会结束
		// fmt.Println(i)
		fmt.Printf("Get %v, Count: %v\n", q.Get(), i+1)
	}
	fmt.Println("ad")

	t2 := time.Now()

	fmt.Println(t2.Sub(t1))
}
