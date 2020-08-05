package main

import "testing"

var caseOne = -123
var caseOneExpect = -321

var (
	caseTwo       = 1534236469
	caseTwoExpect = 0
)

func Test_reverse(t *testing.T) {
	expect := reverse(caseOne)
	if expect != caseOneExpect {
		t.Errorf("整数反转 测试用例:%d 错误结果:%d", caseOne, expect)
	}
	if expect := reverse(caseTwo); expect != caseTwoExpect {
		t.Errorf("整数反转 测试用例:%d 错误结果:%d", caseTwo, expect)
	}

}
