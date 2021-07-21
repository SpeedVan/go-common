package queue

import "fmt"

type SPSlot struct {
	wbusy uint32 // 写行为相互抢占用
	rbusy uint32 // 读行为相互抢占
	ready bool   // 读写之间临界值
	val   interface{}
	next  *SPSlot
}

type SPSlotArr []SPSlot

func NewSPSlotArr(cap uint32) *SPSlotArr {
	arr := make(SPSlotArr, cap)
	for i := uint32(0); i < cap-1; i++ {
		fmt.Printf("arr[%v]:%p, next:%p\n", i, &arr[i], &arr[i+1])
		arr[i].next = &arr[i+1]
	}

	return &arr
}

func NewSPSlotRing(cap uint32) *SPSlotArr {
	arr := NewSPSlotArr(cap)
	fmt.Printf("arr[%v]:%p, next:%p\n", cap-1, &(*arr)[cap-1], &(*arr)[0])
	(*arr)[cap-1].next = &(*arr)[0]
	return arr
}

func (s *SPSlotArr) HeadSlot() *SPSlot {
	return &(*s)[0]
}

func (s *SPSlotArr) LastSlot() *SPSlot {
	return &(*s)[len(*s)-1]
}
