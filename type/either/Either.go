package either

// Either todo
type Either interface {
	// Value()
}

// LeftValue todo
type LeftValue struct {
	value interface{}
}

// Left todo
func Left(value interface{}) *LeftValue {
	return &LeftValue{value: value}
}

// Value todo
func (s *LeftValue) Value() interface{} {
	return s.value
}

// RightValue todo
type RightValue struct {
	value interface{}
}

// Right todo
func Right(value interface{}) *RightValue {
	return &RightValue{value: value}
}

// Value todo
func (s *RightValue) Value() interface{} {
	return s.value
}
