package leetcode_hot_100

// TODO  again
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxCount := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if maxCount < dp[i] {
			maxCount = dp[i]
		}
	}
	return maxCount
}
