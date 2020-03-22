package main

import "math"

func main() {

}

// 只允许交易一次的情况下
// 贪心算法
func maxProfit(prices []int) int {
	// 记录当前最大回报
	mp := 0
	// low 表示前 N 天的最低的价位
	low := math.MaxInt32
	// 因为 low 是从左往右遍历，所以能保证当前 price 至少是在 low 之后出现的。
	for _, price := range prices {
		if price < low {
			low = price
		}
		mp = max(mp, price-low)
	}
	return mp
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
