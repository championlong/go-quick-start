package ioc

import "fmt"

type TestCase interface {
	DoTest() bool
}

type JunitApplication struct {
	TestCases []TestCase
}

func NewJunitApplication() *JunitApplication {
	return &JunitApplication{
		TestCases: make([]TestCase, 0),
	}
}

func (self *JunitApplication) Register(testCase TestCase) {
	self.TestCases = append(self.TestCases, testCase)
}

func (self *JunitApplication) Run() {
	for _, testCase := range self.TestCases{
		result := testCase.DoTest()
		if result {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
}

type UserServiceTest struct {
}

func (self *UserServiceTest) DoTest() bool {
	fmt.Println("test")
	return true
}
