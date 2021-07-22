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

type BaseSlot struct {
	Slot
	_wbusy uint32 // 写行为相互抢占用
	_rbusy uint32 // 读行为相互抢占
	_ready bool   // 读写之间临界值
	_next  Slot
}

func (s *BaseSlot) wBusy() *uint32 {
	return &s._wbusy
}

func (s *BaseSlot) rBusy() *uint32 {
	return &s._rbusy
}

func (s *BaseSlot) ready() *bool {
	return &s._ready
}

func (s *BaseSlot) next() Slot {
	return s._next
}

func (s *BaseSlot) setNext(n Slot) {
	s._next = n
}

type DefaultSlotArr struct {
	arr []*Slot
}

func (s *DefaultSlotArr) HeadSlot() Slot {
	return *s.arr[0]
}

func (s *DefaultSlotArr) LastSlot() Slot {
	return *s.arr[len(s.arr)-1]
}

func (s *DefaultSlotArr) Len() uint32 {
	return uint32(len(s.arr))
}

func (s *DefaultSlotArr) Get(idx uint32) Slot {
	return *s.arr[idx]
}

func NewDefaultSlotArr(arr []*Slot) SlotArr {
	return &DefaultSlotArr{arr}
}
