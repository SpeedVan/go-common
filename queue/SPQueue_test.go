package queue

import (
	"fmt"
	"testing"
	"time"
)

type MySlot struct {
	BaseSlot
	t        time.Time
	abstract string
	action   string
	data     []byte
}

type MySlotArr []MySlot

func NewMySlotArr(cap uint32) *MySlotArr {
	arr := make(MySlotArr, cap)
	for i := uint32(0); i < cap-1; i++ {
		fmt.Printf("arr[%v]:%p, next:%p\n", i, &arr[i], &arr[i+1])
		arr[i].setNext(&arr[i+1])
	}

	return &arr
}

func NewMySlotRing(cap uint32) *MySlotArr {
	arr := NewMySlotArr(cap)
	fmt.Printf("arr[%v]:%p, next:%p\n", cap-1, &(*arr)[cap-1], &(*arr)[0])
	(*arr)[cap-1].setNext(&(*arr)[0])
	return arr
}

func (s *MySlotArr) HeadSlot() Slot {
	return &(*s)[0]
}

func (s *MySlotArr) LastSlot() Slot {
	return &(*s)[len(*s)-1]
}

func (s *MySlotArr) Len() uint32 {
	return uint32(len(*s))
}

func (s *MySlotArr) Get(idx uint32) Slot {
	return &(*s)[idx]
}

func TestMyQueue(t *testing.T) {
	q := NewQueue(func() SlotArr {
		return NewSlotRing(func() SlotArr {
			return NewMySlotRing(1024)
		})
	})
	// m := map[interface{}]int{}
	// for i := 0; i < 10000; i++ {
	// 	m[i] = i
	// }
	// ch := make(chan interface{})

	routine1 := func() {
		publisher := q.Publisher("p1")
		for i := 0; i < 10000; i = i + 2 {
			publisher.Put(func(s Slot) { (s.(*MySlot)).t = time.Now() })
		}
	}

	routine2 := func() {
		publisher := q.Publisher("p2")
		for i := 1; i < 10000; i = i + 2 {
			publisher.Put(func(s Slot) { (s.(*MySlot)).t = time.Now() })
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
		subscribe.Get(func(s Slot) {})
		// ch <- subscribe.Get()
		fmt.Println(i)
		// fmt.Printf("Get %v, Count: %v\n", q.Get(), i+1)
	}
	fmt.Println("end")

	t2 := time.Now()

	fmt.Println(t2.Sub(t1))
	// fmt.Scanln()
}
