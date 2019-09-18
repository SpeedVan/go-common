package maybe

type MaybeInt interface {
	Maybe
}

type Integer int

func (s Integer) maybe() bool {
	return true
}
