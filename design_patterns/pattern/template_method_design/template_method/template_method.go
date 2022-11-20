package template_method

import "fmt"

type ICallback interface {
	methodToCallback()
}

type BClass struct {
}

func (self BClass) process(callback ICallback) {
	callback.methodToCallback()
}

type AClass struct {
}

func (self AClass) methodToCallback() {
	fmt.Println("callback me")
}
