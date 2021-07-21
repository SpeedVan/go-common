package snapshot

import (
	"fmt"
	"time"

	"github.com/alpha-supsys/go-common/simple"
)

type canString interface {
	String() string
}

type Str struct {
	str string
}

func (s *Str) String() string {
	return s.str
}

type Event struct {
	t        time.Time
	abstract string
	data     canString
}

type Snapshot struct {
	abstract string
	data     interface{}
	queue    chan canString
	hander   func(*Event, interface{}, string) (interface{}, error)
}

func NewSnapshot(init interface{}, hander func(*Event, interface{}, string) (interface{}, error)) *Snapshot {
	return &Snapshot{
		abstract: simple.RandString(16),
		data:     init,
		queue:    make(chan canString),
		hander:   hander,
	}
}

func (s *Snapshot) Get() *Event {
	d := <-s.queue
	fmt.Println()
	return &Event{
		t:        time.Now(),
		abstract: simple.Sha1(s.abstract + d.String()),
		data:     d,
	}
}

func (s *Snapshot) Put(v canString) {
	s.queue <- v
}

func (s *Snapshot) PutStr(v string) {
	s.queue <- &Str{
		str: v,
	}
}

func (s *Snapshot) Run() {
	go func() {
		for {
			e := s.Get()
			d, err := s.hander(e, s.data, s.abstract)
			if err != nil {
				return
			}
			s.data = d
			s.abstract = e.abstract
		}
	}()
}
