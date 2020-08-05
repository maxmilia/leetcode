package mergesortedarray

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	result := []int{1, 2, 2, 3, 5, 6}
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	case2nums1 := []int{0}
	case2nums2 := []int{1}
	case2m := 0
	case2n := 1
	case2result := []int{1}
	merge(nums1, m, nums2, n)
	merge(case2nums1, case2m, case2nums2, case2n)
	if !reflect.DeepEqual(nums1, result) {
		t.Error("测试未通过")
	}
	if !reflect.DeepEqual(case2nums1, case2result) {
		t.Error("测试用例2未通过")
	}
}
