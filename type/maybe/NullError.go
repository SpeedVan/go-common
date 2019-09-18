package maybe

type nullError struct {
	Maybe
	error
}

func (s *nullError) Error() string {
	return "null error"
}

func (s *nullError) maybe() bool {
	return false
}

var NullError = &nullError{}

func Error() *nullError {
	return NullError
}
