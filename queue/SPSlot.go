package queue

import "fmt"

type SPSlot struct {
	BaseSlot
	_val interface{}
}

type SPSlotArr []SPSlot

func NewSPSlotArr(cap uint32) *SPSlotArr {
	arr := make(SPSlotArr, cap)
	for i := uint32(0); i < cap-1; i++ {
		fmt.Printf("arr[%v]:%p, next:%p\n", i, &arr[i], &arr[i+1])
		arr[i].setNext(&arr[i+1])
	}

	return &arr
}

func NewSPSlotRing(cap uint32) *SPSlotArr {
	arr := NewSPSlotArr(cap)
	fmt.Printf("arr[%v]:%p, next:%p\n", cap-1, &(*arr)[cap-1], &(*arr)[0])
	(*arr)[cap-1].setNext(&(*arr)[0])
	return arr
}

func (s *SPSlotArr) HeadSlot() *SPSlot {
	return &(*s)[0]
}

func (s *SPSlotArr) LastSlot() *SPSlot {
	return &(*s)[len(*s)-1]
}

func (s *SPSlotArr) Len() uint32 {
	return uint32(len(*s))
}
func (s *SPSlotArr) Get(i uint32) Slot {
	return &((*s)[i])
}
