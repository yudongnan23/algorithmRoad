package leetcode_hot_100

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) <= 1 {
		return maxProfit
	}

	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		curProfit := prices[i] - minPrice
		if curProfit > maxProfit {
			maxProfit = curProfit
		}
	}
	return maxProfit
}
