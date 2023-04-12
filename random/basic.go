package random

import (
	"reflect"

	fake "github.com/brianvoe/gofakeit/v6"
)

type ItemBuilder[T any] struct {
	value T
}

// following types will be fill with random value without customization:
//  - struct (only exported fields)
//  - int
//  - int8
//  - int16
//  - int32
//  - int64
//  - uint
//  - uint8
//  - uint16
//  - uint32
//  - uint64
//  - bool
//  - float32
//  - float64
//  - string
func New[T any]() ItemBuilder[T] {
	t := simpleRandom[T]()

	return ItemBuilder[T]{
		value: t,
	}
}

func (builder ItemBuilder[T]) Set(f func(*T)) ItemBuilder[T] {
	f(&builder.value)
	return builder
}

func (builder ItemBuilder[T]) Build() T {
	return builder.value
}

func simpleRandom[T any]() T {
	res := new(T)
	rv := reflect.ValueOf(res)
	setRandomValue(rv.Elem())
	return *res
}

func setRandomValue(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if field.CanSet() {
				setRandomValue(field)
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			setRandomValue(rv.Index(i))
		}
	default:
		rv.Set(reflect.ValueOf(fakeValue(rv.Type())))
	}
}

func fakeValue(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Int8:
		return fake.Int8()
	case reflect.Int:
		return int(fake.Int16())
	case reflect.Int16:
		return fake.Int16()
	case reflect.Int32:
		return fake.Int32()
	case reflect.Int64:
		return fake.Int64()
	case reflect.Uint:
		return uint(fake.Uint16())
	case reflect.Uint16:
		return fake.Uint16()
	case reflect.Uint8:
		return fake.Uint8()
	case reflect.Uint32:
		return fake.Uint32()
	case reflect.Uint64:
		return fake.Uint64()
	case reflect.Bool:
		return fake.Bool()
	case reflect.Float32:
		return fake.Float32()
	case reflect.Float64:
		return fake.Float64()
	case reflect.String:
		return fake.SentenceSimple()
	case reflect.Struct:
		return fakeStruct(t)
	case reflect.Ptr:
		return reflect.New(t.Elem()).Interface()
	default:
		return nil
	}
}

func fakeStruct(t reflect.Type) interface{} {
	return reflect.New(t).Elem().Interface()
}

type ItemsBuilder[T any] struct {
	value []T
}

func (m ItemsBuilder[T]) Build() []T {
	return m.value
}

func NewSlice[T any](builderDelegate func() Builder[T], amount uint8) ItemsBuilder[T] {
	value := make([]T, amount)
	for i := uint8(0); i < amount; i++ {
		value[i] = builderDelegate().Build()
	}
	return ItemsBuilder[T]{
		value: value,
	}
}
