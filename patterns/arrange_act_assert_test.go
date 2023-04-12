package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAAA(t *testing.T) {
	assert := assert.New(t)
	sampleAddFunc := func(a, b int) int { return a + b }

	/* ---------------------------- wrapped AAA test ---------------------------- */
	RunWrappedAAATest(func() Act {
		const num1 = 7
		const num2 = 8
		return func() Assert {
			actual := sampleAddFunc(num1, num2)
			return func() {
				const expected = 15
				assert.Equal(expected, actual)
			}
		}
	})

	/* -------------------------------- AAA test -------------------------------- */
	type arrangeRet = struct {
		num1 int
		num2 int
	}
	RunAAATest(
		func() arrangeRet {
			return arrangeRet{num1: 7, num2: 8}
		},
		func(arr arrangeRet) int {
			return sampleAddFunc(arr.num1, arr.num2)
		},
		func(_ arrangeRet, actual int) {
			const expected = 15
			assert.Equal(expected, actual)
		})

	/* ---------------------------- chained AAA test ---------------------------- */
	RunChainedAAATest[arrangeRet, int]()(func() arrangeRet {
		return arrangeRet{num1: 7, num2: 8}
	})(func(arr arrangeRet) int {
		return sampleAddFunc(arr.num1, arr.num2)
	})(func(_ arrangeRet, actual int) {
		const expected = 15
		assert.Equal(expected, actual)
	})

	/* --------------------------------- struct --------------------------------- */
	{
		var num1, num2 int
		NewAAA[int]().Given(func() {
			num1 = 7
			num2 = 8
		}).When(func() int {
			return sampleAddFunc(num1, num2)
		}).Then(func(actual int) {
			const expected = 15
			assert.Equal(expected, actual)
		})
	}
}
