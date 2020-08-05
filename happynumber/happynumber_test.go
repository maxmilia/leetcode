package happynumber

import "testing"

func Test_isHappy(t *testing.T) {
	type testCase struct {
		Input  int
		Expect bool
	}
	var testGroup = map[string]testCase{
		"1": testCase{Input: 19, Expect: true},
	}
	for k, v := range testGroup {
		t.Run(k, func(t *testing.T) {
			g := isHappy(v.Input)
			if g != v.Expect {
				t.Errorf("%s失败", k)
			}
		})
	}
}
