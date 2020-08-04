package main

// 88 合并两个有序数组

func merge(nums1 []int, m int, nums2 []int, n int) {
	oneIndex := m - 1
	twoIndex := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if twoIndex < 0 {
			break
		}
		if oneIndex >= 0 && nums1[oneIndex] > nums2[twoIndex] {
			nums1[i] = nums1[oneIndex]
			oneIndex--
		} else {
			nums1[i] = nums2[twoIndex]
			twoIndex--
		}
	}
}
