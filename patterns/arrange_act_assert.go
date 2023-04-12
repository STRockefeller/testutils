package patterns

/* -------------------------------------------------------------------------- */
/*                                 Wrapped AAA                                */
/* -------------------------------------------------------------------------- */

type Assert func()
type Act func() Assert
type Arrange func() Act

func RunWrappedAAATest(aaa Arrange) {
	aaa()()()
}

/* -------------------------------------------------------------------------- */
/*                                     AAA                                    */
/* -------------------------------------------------------------------------- */

func RunAAATest[T1, T2 any](arrange func() T1, act func(T1) T2, assert func(T1, T2)) {
	arrangeReturn := arrange()
	assert(arrangeReturn, act(arrangeReturn))
}

/* -------------------------------------------------------------------------- */
/*                                 Chained AAA                                */
/* -------------------------------------------------------------------------- */

func RunChainedAAATest[T1, T2 any]() func(func() T1) func(func(T1) T2) func(func(T1, T2)) {
	return func(arrange func() T1) func(func(T1) T2) func(func(T1, T2)) {
		return func(act func(T1) T2) func(func(T1, T2)) {
			return func(assert func(T1, T2)) {
				RunAAATest(arrange, act, assert)
			}
		}
	}
}

/* -------------------------------------------------------------------------- */
/*                                  Structure                                 */
/* -------------------------------------------------------------------------- */

// AAA struct represents the test context and provides methods to set up the
// preconditions, perform the action, and make assertions.
type AAA[T any] struct {
	arrange func()
	act     func() T
	assert  func(actual T)
}

func NewAAA[T any]() *AAA[T] {
	return new(AAA[T])
}

// Given sets up the Arrange function with the given setup function.
func (a *AAA[T]) Given(setup func()) *AAA[T] {
	a.arrange = setup
	return a
}

// When sets up the Act function with the given action function.
func (a *AAA[T]) When(action func() T) *AAA[T] {
	a.act = action
	return a
}

// Then sets up the Assert function with the given assertion function.
func (a *AAA[T]) Then(assertion func(actual T)) {
	a.assert = assertion
	a.run()
}

// alias of Given
func (a *AAA[T]) Arrange(setup func()) *AAA[T] {
	return a.Given(setup)
}

// alias of When
func (a *AAA[T]) Act(action func() T) *AAA[T] {
	return a.When(action)
}

// alias of Then
func (a *AAA[T]) Assert(assertion func(actual T)) {
	a.Then(assertion)
}

// run executes the test by calling the Arrange, Act, and Assert functions
// in the proper order.
func (a *AAA[T]) run() {
	if a.arrange != nil {
		a.arrange()
	}
	if a.act != nil {
		actual := a.act()
		if a.assert != nil {
			a.assert(actual)
		}
	}
}
