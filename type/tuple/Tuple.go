package tuple

type Tuple struct {
	V1 interface{}
	V2 interface{}
}

func (s *Tuple) V() (interface{}, interface{}) {
	return s.V1, s.V2
}
