package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	low := 1<<31 - 1
	high := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		} else if prices[i]-low > high {
			high = prices[i] - low
		}

	}
	return high
}
