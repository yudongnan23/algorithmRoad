package leetcode_hot_100

// TODO again
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum / 2
	if target*2 != sum {
		return false
	}
	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	// fmt.Println(dp)
	return dp[target] == target
}
