package leetcode_hot_100

// TODO three
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = i + 1
		for _, coin := range coins {
			if i < coin {
				continue
			}
			curMaxCoins := dp[i-coin] + 1
			if dp[i-coin] < i-coin+1 && curMaxCoins < dp[i] {
				dp[i] = curMaxCoins
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
