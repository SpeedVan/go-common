package maybe

type MaybeInt interface {
	Maybe
}

type Integer int

func (s Integer) maybe() bool {
	return true
}

func (s Integer) Just(func(interface{})) {

}
func (s Integer) Nothing(func()) {

}
