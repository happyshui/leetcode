package leetcode

func maxProfit(prices []int) int {
	length := len(prices)
	min := 20000 // define int max value
	var maxsum int
	for i := 0; i < length; i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > maxsum {
			maxsum = prices[i] - min
		}
	}
	return maxsum
}
