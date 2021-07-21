package queue

import "fmt"

type Slot interface {
	wBusy() *uint32
	rBusy() *uint32
	ready() *bool
	next() Slot
	setNext(Slot)
}

type SlotArr interface {
	HeadSlot() Slot
	LastSlot() Slot
	Len() uint32
	Get(uint32) Slot
}

func NewSlotArr(arrCreator func() SlotArr) SlotArr {
	arr := arrCreator()
	cap := arr.Len()
	for i := uint32(0); i < cap-1; i++ {
		fmt.Printf("arr[%v]:%p, next:%p\n", i, arr.Get(i), arr.Get(i+1))
		arr.Get(i).setNext(arr.Get(i + 1))
	}
	return arr
}

func NewSlotRing(arrCreator func() SlotArr) SlotArr {
	arr := NewSlotArr(arrCreator)
	arr.Get(arr.Len() - 1).setNext(arr.Get(0))
	return arr
}
