package ioc

import "fmt"

type TestCase interface {
	doTest() bool
}

func run() {
	if doTest() {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}
}

func doTest() bool {
	return true
}

type JunitApplication struct {
	testCases []TestCase
}

func NewJunitApplication() *JunitApplication {
	return &JunitApplication{
		testCases: make([]TestCase, 0),
	}
}

func (self *JunitApplication) Register(testCase TestCase) {
	self.testCases = append(self.testCases, testCase)
}

type UserServiceTest struct {
}

func (self *UserServiceTest) doTest() bool {
	return true
}
