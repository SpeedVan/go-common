package queue

type Slot interface {
	Ready() bool
	Val() interface{}
	Next() Slot
}

type DSlot struct {
	ready bool
	val   interface{}
	next  *DSlot
}

func (s *DSlot) Ready() bool {
	return s.ready
}

func (s *DSlot) Val() interface{} {
	return s.val
}

func (s *DSlot) Next() Slot {
	return s.next
}

type DSlotArr []DSlot

func NewDSlotArr(cap uint32) *DSlotArr {
	arr := make(DSlotArr, cap)
	for i := uint32(0); i < cap-1; i++ {
		// fmt.Printf("arr[%v]:%p, next:%p\n", i, &arr[i], &arr[i+1])
		arr[i].next = &arr[i+1]
	}

	return &arr
}

func NewDSlotRing(cap uint32) *DSlotArr {
	arr := NewDSlotArr(cap)
	// fmt.Printf("arr[%v]:%p, next:%p\n", cap-1, &(*arr)[cap-1], &(*arr)[0])
	(*arr)[cap-1].next = &(*arr)[0]
	return arr
}

func (s *DSlotArr) HeadSlot() *DSlot {
	return &(*s)[0]
}

func (s *DSlotArr) LastSlot() *DSlot {
	return &(*s)[len(*s)-1]
}
