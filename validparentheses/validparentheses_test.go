package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	type TestCase struct {
		Input  string
		Expect bool
	}
	t1 := TestCase{Input: "]", Expect: false}
	o := isValid(t1.Input)
	if o != t1.Expect {
		t.Errorf("有效的括号")
	}
}
