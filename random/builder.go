package random

type Builder[T any] interface {
	Build() T
}
