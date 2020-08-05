package besttimetobuyandsellstock

import "testing"

type testCase struct {
	Input  []int
	Expect int
}

func Test_maxProfit(t *testing.T) {
	testGroup := map[string]testCase{
		"1": testCase{Input: []int{7, 1, 5, 3, 6, 4}, Expect: 5},
		"2": testCase{Input: []int{7, 6, 4, 3, 1}, Expect: 0},
	}

	for k, v := range testGroup {
		t.Run(k, func(t *testing.T) {
			got := maxProfit(v.Input)
			if v.Expect != got {
				t.Errorf("%s失败", k)
			}
		})
	}
}
