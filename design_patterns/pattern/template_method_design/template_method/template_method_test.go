package template_method

import "testing"

func TestCallbackFunc(t *testing.T) {
	var bClass BClass
	bClass.process(new(AClass))
}
