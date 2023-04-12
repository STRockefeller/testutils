# TestUtils - Golang Library

TestUtils is a Golang library that provides utility functions for testing purposes. It includes features such as constraint building and suite setup/teardown functions, which can be used in test suites to simplify and streamline the testing process.

## Installation

To use TestUtils in your Golang project, you need to install it using the `go get` command:

```bash
go get github.com/STRockefeller/testutils
```

## Features

TestUtils includes the following features:

### Constraint Builder

The `ConstraintBuilder` allows you to build constraints for generating random values that meet certain conditions. It provides the following functions:

```go
type ConstraintBuilder[T any] struct {
    conditions []func(T) bool
}

func NewConstraintBuilder[T any]() ConstraintBuilder[T]
func (b ConstraintBuilder[T]) Build() T
func (b ConstraintBuilder[T]) WithConditions(conditions ...func(T) bool) ConstraintBuilder[T]
```

Example usage:

```go
import (
	"github.com/STRockefeller/testutils/random"
	"github.com/STRockefeller/go-linq"
)

func TestRandomNumber(t *testing.T) {
    cb := random.NewConstraintBuilder[int]()
    cb.WithConditions(func(n int) bool { return n > 0 }).WithConditions(func(n int) bool { return n < 100 })
    randomNumber := cb.Build()
    // Use the randomNumber in your test
}
```

### OtherBuilders

- ItemBuilder: give you random value of specified type
- ItemsBuilder: give you random slice value of specified type

### Suite Template

The `SuiteTemplate` provides a template for creating test suites with setup and teardown functions. It includes the following functions:

```go
type SuiteTemplate struct {
    // ...
}

func NewSuiteTemplate() *SuiteTemplate
func (s *SuiteTemplate) SetBeforeTestFuncs(funcs ...beforeFunc) *SuiteTemplate
func (s *SuiteTemplate) AddBeforeTestFuncs(funcs ...beforeFunc) *SuiteTemplate
func (s *SuiteTemplate) SetSetupTestFuncs(funcs ...setupFunc) *SuiteTemplate
func (s *SuiteTemplate) AddSetupTestFuncs(funcs ...setupFunc) *SuiteTemplate
func (s *SuiteTemplate) SetSetupAllFuncs(funcs ...setupFunc) *SuiteTemplate
func (s *SuiteTemplate) AddSetupAllFuncs(funcs ...setupFunc) *SuiteTemplate
func (s *SuiteTemplate) SetSetupSubTestFuncs(funcs ...setupFunc) *SuiteTemplate
func (s *SuiteTemplate) AddSetupSubTestFuncs(funcs ...setupFunc) *SuiteTemplate
func (s *SuiteTemplate) SetupSuite()
func (s *SuiteTemplate) TearDownSuite()
func (s *SuiteTemplate) SetupTest()
func (s *SuiteTemplate) TearDownTest()
func (s *SuiteTemplate) SetupSubTest()
func (s *SuiteTemplate) TearDownSubTest()
func (s *SuiteTemplate) BeforeTest(suiteName, testName string)
```

Example usage:

```go
package main

import (
	"github.com/yourusername/testutils/suite"
	"testing"
)

func TestMySuite(t *testing.T) {
	// Create a new SuiteTemplate
	suiteTemplate := suite.NewSuiteTemplate()

	// Set up before test functions
	suiteTemplate.SetBeforeTestFuncs(
		func(suiteName, testName string) func(string, string) {
			return func(suiteName, testName string) {
				// Do something before each test
			}
		},
	)

	// Set up setup test functions
	suiteTemplate.SetSetupTestFuncs(
		func() func() {
			return func() {
				// Do something before each test
			}
		},
	)

	// Set up setup all functions
	suiteTemplate.SetSetupAllFuncs(
		func() func() {
			return func() {
				// Do something before the entire suite
			}
		},
	)

	// Set up setup sub-test functions
	suiteTemplate.SetSetupSubTestFuncs(
		func() func() {
			return func() {
				// Do something before each sub-test
			}
		},
	)

	// Create a testify suite using the SuiteTemplate
	testSuite := suite.New(suiteTemplate)

	// Add tests to the suite
	testSuite.Run(t, new(MySuite))
}

```

License
testutils is released under the [MIT License](./LICENSE).

Contact
For any questions, issues, or feedback, please open an issue on the [GitHub repository](github.com/STRockefeller/testutils).
