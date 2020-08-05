package main

// 32位数字，有符号的
// 溢出为0

func reverse(x int) int {
	const MIN = -(1 << 31)
	const MAX = 1<<31 - 1
	r := 0
	for x != 0 {
		r = r*10 + x%10
		if r <= MIN || r >= MAX {
			return 0
		}
		x = x / 10
	}
	return r
}
