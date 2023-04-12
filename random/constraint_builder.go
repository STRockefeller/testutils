package random

import "github.com/STRockefeller/go-linq"

type ConstraintBuilder[T any] struct {
	conditions []func(T) bool
}

func NewConstraintBuilder[T any]() ConstraintBuilder[T] {
	return ConstraintBuilder[T]{}
}

func (b ConstraintBuilder[T]) Build() T {
	var temp T
	for {
		temp = New[T]().Build()
		if linq.NewLinq(b.conditions).All(func(f func(T) bool) bool { return f(temp) }) {
			return temp
		}
	}
}

func (b ConstraintBuilder[T]) WithConditions(conditions ...func(T) bool) ConstraintBuilder[T] {
	b.conditions = conditions
	return b
}
