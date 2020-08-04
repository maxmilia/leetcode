package main

import (
	"testing"
)

func Test_climbStairs(t *testing.T) {
	n := 3
	s := climbStairs(n)
	if s != 3 {
		t.Errorf("测试未通过：爬楼梯方法数为 %d", s)
	}
}
func Test_climbStairsUseDp(t *testing.T) {
	n := 3
	s := climbStairsUseDp(n)
	if s != 3 {
		t.Errorf("测试未通过：爬楼梯方法数为 %d", s)
	}
}
