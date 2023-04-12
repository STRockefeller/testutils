package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	assert.NotZero(New[int]().Build())
	assert.NotZero(New[string]().Build())
	assert.NotZero(New[struct{ A string }]())
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	// Test setting a custom value using Set method
	assert.Equal(42, New[int]().Set(func(i *int) {
		*i = 42
	}).Build())
}

func TestNewSlice(t *testing.T) {
	assert := assert.New(t)
	// Test creating a new slice of int type
	builderDelegate := func() Builder[int] {
		return New[int]()
	}
	amount := uint8(5)
	multiBuilder := NewSlice(builderDelegate, amount)
	items := multiBuilder.Build()
	assert.Len(items, int(amount))

	// Test creating a new slice of string type
	builderDelegateStr := func() Builder[string] {
		return New[string]()
	}
	amountStr := uint8(3)
	multiBuilderStr := NewSlice(builderDelegateStr, amountStr)
	itemsStr := multiBuilderStr.Build()
	assert.Len(itemsStr, int(amountStr))
}

func TestCondition(t *testing.T) {
	assert := assert.New(t)
	type user struct {
		Name string
		Age  int
	}
	nameShouldBeShort := func(u user) bool { return len(u.Name) <= 20 }
	userShouldBeAdult := func(u user) bool { return u.Age >= 18 }
	noMonsters := func(u user) bool { return u.Age <= 150 }
	u := NewConstraintBuilder[user]().WithConditions(nameShouldBeShort, userShouldBeAdult, noMonsters).Build()
	assert.True(nameShouldBeShort(u))
	assert.True(userShouldBeAdult(u))
	assert.True(noMonsters(u))
}
