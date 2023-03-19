package to_offer

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minCost := prices[0]
	maxProfit := 0
	for i := 1; i <= len(prices); i++ {
		curProfit := prices[i] - minCost

		if curProfit > maxProfit {
			maxProfit = curProfit
		}

		if prices[i] < minCost {
			minCost = prices[i]
		}
	}

	return maxProfit
}
