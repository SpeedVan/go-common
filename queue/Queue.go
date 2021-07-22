package queue

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type Queue interface {
	Publisher(id string) Publisher
	Subscribe(id string) Subscribe
}

type Publisher interface {
	Put(handler func(Slot))
}

type Subscribe interface {
	Get(handler func(Slot))
}

type BaseQueue struct {
	slots []SlotArr

	cap   uint32
	gSlot Slot
	pSlot Slot

	scaleChan chan int
	stopChan  chan int

	slotArrCreator func() SlotArr
}

func NewQueue(slotArrCreator func() SlotArr) Queue {
	arr := slotArrCreator()

	q := &BaseQueue{
		slots:          []SlotArr{arr},
		cap:            arr.Len(),
		gSlot:          arr.HeadSlot(),
		pSlot:          arr.HeadSlot(),
		scaleChan:      make(chan int),
		stopChan:       make(chan int),
		slotArrCreator: slotArrCreator,
	}
	// q.asyncAutoscale()
	return q
}

func (s *BaseQueue) asyncAutoscale() {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			case <-s.scaleChan:
			}
			if len(s.scaleChan) > 5 {
				arr := s.slotArrCreator()
				arr.LastSlot().setNext(s.slots[0].HeadSlot())
				s.slots[len(s.slots)-1].LastSlot().setNext(arr.HeadSlot())
				s.slots = append(s.slots, arr)
				fmt.Println("scaled")
			}
		}
	}()
}

func (s *BaseQueue) Publisher(id string) Publisher {
	return &DefalutPublisher{
		id:    id,
		slot:  s.pSlot, // 起始参考值，
		queue: s,
	}
}

func (s *BaseQueue) Subscribe(id string) Subscribe {
	return &DefaultSubscribe{
		id:    id,
		slot:  s.gSlot,
		queue: s,
	}
}

type DefalutPublisher struct {
	id    string // publisher标识
	slot  Slot   // 每个publisher有自己的offset的意思，用链表则每个都有自己的引用
	queue Queue
}

func (s *DefalutPublisher) Put(handler func(Slot)) {
	for {
		if *s.slot.rBusy() > 0 { // 写追尾读，没必要继续
			runtime.Gosched()
			continue
		}

		if atomic.CompareAndSwapUint32(
			s.slot.wBusy(),
			0,
			1,
		) {
			if *s.slot.ready() { // 抢占成功后，若果当前写过了还没读，则跳过
				*s.slot.wBusy() = 0
				s.slot = s.slot.next()

				continue
			} else {
				for { // 存疑，实际可能并不存在这样的状态
					if !*s.slot.ready() { // 存疑，实际可能并不存在这样的状态
						handler(s.slot)
						*s.slot.ready() = true
						*s.slot.wBusy() = 0
						return
					} else {
						break
					}
				}
			}
		} else {
			s.slot = s.slot.next()
		}

	}
}

type DefaultSubscribe struct {
	id    string // subscribe标识
	slot  Slot   // 每个subscribe有自己的offset的意思，用链表则每个都有自己的引用
	queue Queue
}

func (s *DefaultSubscribe) Get(handler func(Slot)) {
	for {
		if *s.slot.wBusy() > 0 { // 读追写，让对方先前进
			runtime.Gosched()
			continue
		}

		if atomic.CompareAndSwapUint32(
			s.slot.rBusy(),
			0,
			1,
		) {
			if !*s.slot.ready() { // 抢占成功后，若果当前读过了还没写，则跳过
				*s.slot.rBusy() = 0
				s.slot = s.slot.next()
				continue
			} else {
				for { // 存疑，实际可能并不存在这样的状态
					if *s.slot.ready() { // 存疑，实际可能并不存在这样的状态
						handler(s.slot)
						*s.slot.ready() = false
						*s.slot.rBusy() = 0
						return
					} else {
						break
					}
				}
			}
		}

	}
}
