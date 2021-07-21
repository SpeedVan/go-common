package queue

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type SPQueue struct {
	slots []*SPSlotArr

	cap   uint32
	gSlot *SPSlot
	pSlot *SPSlot

	scaleChan chan int
	stopChan  chan int
}

func NewSPQueue(cap uint32) *SPQueue {
	arr := NewSPSlotRing(cap)

	q := &SPQueue{
		slots:     []*SPSlotArr{arr},
		cap:       cap,
		gSlot:     arr.HeadSlot(),
		pSlot:     arr.HeadSlot(),
		scaleChan: make(chan int),
		stopChan:  make(chan int),
	}
	// q.asyncAutoscale()
	return q
}

func (s *SPQueue) asyncAutoscale() {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			case <-s.scaleChan:
			}
			if len(s.scaleChan) > 5 {
				arr := NewSPSlotArr(s.cap)
				arr.LastSlot().next = s.slots[0].HeadSlot()
				s.slots[len(s.slots)-1].LastSlot().next = arr.HeadSlot()
				s.slots = append(s.slots, arr)
				fmt.Println("scaled")
			}
		}
	}()
}

func (s *SPQueue) Publisher(id string) *Publisher {
	return &Publisher{
		id:    id,
		slot:  s.pSlot, // 起始参考值，
		queue: s,
	}
}

func (s *SPQueue) Subscribe(id string) *Subscribe {
	return &Subscribe{
		id:    id,
		slot:  s.gSlot,
		queue: s,
	}
}

type Publisher struct {
	id    string  // publisher标识
	slot  *SPSlot // 每个publisher有自己的offset的意思，用链表则每个都有自己的引用
	queue *SPQueue
}

func (s *Publisher) Put(v interface{}) {
	for {
		if s.slot.rbusy > 0 { // 写追尾读，没必要继续
			runtime.Gosched()
			continue
		}

		if atomic.CompareAndSwapUint32(
			&s.slot.wbusy,
			0,
			1,
		) {
			if s.slot.ready { // 抢占成功后，若果当前写过了还没读，则跳过
				s.slot.wbusy = 0
				s.slot = s.slot.next

				continue
			} else {
				for { // 存疑，实际可能并不存在这样的状态
					if !s.slot.ready { // 存疑，实际可能并不存在这样的状态
						s.slot.val = v
						s.slot.ready = true
						s.slot.wbusy = 0
						return
					} else {
						break
					}
				}
			}
		} else {
			s.slot = s.slot.next
		}

	}
}

type Subscribe struct {
	id    string  // subscribe标识
	slot  *SPSlot // 每个subscribe有自己的offset的意思，用链表则每个都有自己的引用
	queue *SPQueue
}

func (s *Subscribe) Get() interface{} {
	for {
		if s.slot.wbusy > 0 { // 读追写，让对方先前进
			runtime.Gosched()
			continue
		}

		if atomic.CompareAndSwapUint32(
			&s.slot.rbusy,
			0,
			1,
		) {
			if !s.slot.ready { // 抢占成功后，若果当前读过了还没写，则跳过
				s.slot.rbusy = 0
				s.slot = s.slot.next
				continue
			} else {
				for { // 存疑，实际可能并不存在这样的状态
					if s.slot.ready { // 存疑，实际可能并不存在这样的状态
						value := s.slot.val
						s.slot.ready = false
						s.slot.rbusy = 0
						return value
					} else {
						break
					}
				}
			}
		}

	}
}
