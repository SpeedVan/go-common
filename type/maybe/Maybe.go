package maybe

import "github.com/alpha-supsys/go-common/type/either"

type Maybe interface {
	either.Either
	maybe() bool
	Just(func(interface{}))
	Nothing(func())
}
