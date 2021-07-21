package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestSPQueue(t *testing.T) {
	q := NewSPQueue(3)
	// m := map[interface{}]int{}
	// for i := 0; i < 10000; i++ {
	// 	m[i] = i
	// }
	// ch := make(chan interface{})

	routine1 := func() {
		publisher := q.Publisher("p1")
		for i := 0; i < 10000; i = i + 2 {
			publisher.Put(i)
		}
	}

	routine2 := func() {
		publisher := q.Publisher("p2")
		for i := 1; i < 10000; i = i + 2 {
			publisher.Put(i)
		}
	}

	// routine3 := func() {

	// 	for i := 0; i < 10000; i++ {

	// 		delete(m, <-ch)

	// 		if i > 9950 {
	// 			fmt.Println(i)
	// 		}
	// 		if i > 9985 {
	// 			fmt.Println(m)
	// 		}
	// 	}
	// }

	t1 := time.Now()
	go routine1()
	go routine2()
	// go routine3()

	subscribe := q.Subscribe("s1")
	for i := 0; i < 10000; i++ {
		subscribe.Get()
		// ch <- subscribe.Get()
		fmt.Println(i)
		// fmt.Printf("Get %v, Count: %v\n", q.Get(), i+1)
	}
	fmt.Println("end")

	t2 := time.Now()

	fmt.Println(t2.Sub(t1))
	// fmt.Scanln()
}
