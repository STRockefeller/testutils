package suite

import (
	"github.com/STRockefeller/collection"
	"github.com/stretchr/testify/suite"
)

type setupFunc = func() func()
type beforeFunc = func(string, string) func(string, string)

type SuiteTemplate struct {
	suite.Suite
	/* -------------------------- before and after test ------------------------- */
	beforeTestFuncs []beforeFunc
	afterTestFuncs  collection.Stack[func(string, string)]
	/* ------------------------------- setup test ------------------------------- */
	setupTestFuncs    []setupFunc
	teardownTestFuncs collection.Stack[func()]
	/* ----------------------------- setup sub test ----------------------------- */
	setupSubTestFuncs    []setupFunc
	teardownSubTestFuncs collection.Stack[func()]
	/* -------------------------------- setup all ------------------------------- */
	setupAllFuncs    []setupFunc
	teardownAllFuncs collection.Stack[func()]
}

func NewSuiteTemplate() *SuiteTemplate {
	return &SuiteTemplate{
		beforeTestFuncs:      make([]beforeFunc, 0),
		afterTestFuncs:       collection.NewStack[func(string, string)](),
		setupTestFuncs:       make([]setupFunc, 0),
		teardownTestFuncs:    collection.NewStack[func()](),
		setupSubTestFuncs:    make([]setupFunc, 0),
		teardownSubTestFuncs: collection.NewStack[func()](),
		setupAllFuncs:        make([]setupFunc, 0),
		teardownAllFuncs:     collection.NewStack[func()](),
	}
}

func (s *SuiteTemplate) SetBeforeTestFuncs(funcs ...beforeFunc) *SuiteTemplate {
	s.beforeTestFuncs = funcs
	return s
}

func (s *SuiteTemplate) AddBeforeTestFuncs(funcs ...beforeFunc) *SuiteTemplate {
	s.beforeTestFuncs = append(s.beforeTestFuncs, funcs...)
	return s
}

func (s *SuiteTemplate) SetSetupTestFuncs(funcs ...setupFunc) *SuiteTemplate {
	s.setupTestFuncs = funcs
	return s
}

func (s *SuiteTemplate) AddSetupTestFuncs(funcs ...setupFunc) *SuiteTemplate {
	s.setupTestFuncs = append(s.setupTestFuncs, funcs...)
	return s
}

func (s *SuiteTemplate) SetSetupAllFuncs(funcs ...setupFunc) *SuiteTemplate {
	s.setupAllFuncs = funcs
	return s
}

func (s *SuiteTemplate) AddSetupAllFuncs(funcs ...setupFunc) *SuiteTemplate {
	s.setupAllFuncs = append(s.setupAllFuncs, funcs...)
	return s
}

func (s *SuiteTemplate) SetSetupSubTestFuncs(funcs ...setupFunc) *SuiteTemplate {
	s.setupSubTestFuncs = funcs
	return s
}

func (s *SuiteTemplate) AddSetupSubTestFuncs(funcs ...setupFunc) *SuiteTemplate {
	s.setupSubTestFuncs = append(s.setupSubTestFuncs, funcs...)
	return s
}

func parseSetupFuncs(source []setupFunc, tearDownStack *collection.Stack[func()]) {
	for i := len(source) - 1; i >= 0; i-- {
		teardownFunc := source[i]()
		tearDownStack.Push(teardownFunc)
	}
}

func tearDownSetupFuncs(tearDownStack *collection.Stack[func()]) {
	for !tearDownStack.IsEmpty() {
		teardownFunc := tearDownStack.Pop()
		teardownFunc()
	}
}

func (s *SuiteTemplate) SetupSuite() {
	parseSetupFuncs(s.setupAllFuncs, &s.teardownAllFuncs)
}

func (s *SuiteTemplate) TearDownSuite() {
	tearDownSetupFuncs(&s.teardownAllFuncs)
}

func (s *SuiteTemplate) SetupTest() {
	parseSetupFuncs(s.setupTestFuncs, &s.teardownTestFuncs)
}

func (s *SuiteTemplate) TearDownTest() {
	tearDownSetupFuncs(&s.teardownTestFuncs)
}

func (s *SuiteTemplate) SetupSubTest() {
	parseSetupFuncs(s.setupSubTestFuncs, &s.teardownSubTestFuncs)
}

func (s *SuiteTemplate) TearDownSubTest() {
	tearDownSetupFuncs(&s.teardownSubTestFuncs)
}

func (s *SuiteTemplate) BeforeTest(suiteName, testName string) {
	for i := len(s.beforeTestFuncs) - 1; i >= 0; i-- {
		teardownFunc := s.beforeTestFuncs[i](suiteName, testName)
		s.afterTestFuncs.Push(teardownFunc)
	}
}

func (s *SuiteTemplate) AfterTest(suiteName, testName string) {
	for !s.afterTestFuncs.IsEmpty() {
		teardownFunc := s.afterTestFuncs.Pop()
		teardownFunc(suiteName, testName)
	}
}
