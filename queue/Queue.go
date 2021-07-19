package queue

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type Queue struct {
	slots []*DSlotArr

	cap   uint32
	gSlot *DSlot
	pSlot *DSlot

	scaleChan chan int
	stopChan  chan int
}

func NewQueue(cap uint32) *Queue {
	arr := NewDSlotRing(cap)

	q := &Queue{
		slots:     []*DSlotArr{arr},
		cap:       cap,
		gSlot:     arr.HeadSlot(),
		pSlot:     arr.HeadSlot(),
		scaleChan: make(chan int),
		stopChan:  make(chan int),
	}
	q.asyncAutoscale()
	return q
}

func (s *Queue) asyncAutoscale() {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			case <-s.scaleChan:
			}
			if len(s.scaleChan) > 5 {
				arr := NewDSlotArr(s.cap)
				arr.LastSlot().next = s.slots[0].HeadSlot()
				s.slots[len(s.slots)-1].LastSlot().next = arr.HeadSlot()
				s.slots = append(s.slots, arr)
				fmt.Println("scaled")
			}
		}
	}()
}

func (s *Queue) Get() interface{} {
	var gSlot *DSlot
	for {
		gSlot = s.gSlot
		newGSlot := gSlot.next
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.gSlot)),
			unsafe.Pointer(gSlot),
			unsafe.Pointer(newGSlot),
		) {
			for {
				if newGSlot.ready {
					newGSlot.ready = false
					return newGSlot.val
				} else {
					runtime.Gosched()
				}
			}
		} else {
			runtime.Gosched()
		}

	}
}

func (s *Queue) Put(v interface{}) {
	var pSlot *DSlot
	for {
		pSlot = s.pSlot
		newPSlot := pSlot.next
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.pSlot)),
			unsafe.Pointer(pSlot),
			unsafe.Pointer(newPSlot),
		) {
			// fmt.Printf("Put:curr%p next:%p\n", pSlot, newPSlot)
			for {
				if !newPSlot.ready {
					newPSlot.val = v
					newPSlot.ready = true
					return
				} else {
					runtime.Gosched()
				}
			}
		} else {
			runtime.Gosched()
		}
	}
}
