package leetcode_hot_100

// TODO three - 分割等和子集
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}

	halfSum := sum / 2

	if sum%2 != 0 {
		return false
	}

	dp := make([]int, halfSum+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == halfSum {
			return true
		}
		for j := halfSum; j >= nums[i]; j-- {
			curNum := dp[j-nums[i]] + nums[i]
			if curNum == halfSum {
				return true
			}

			if curNum > dp[j] {
				dp[j] = curNum
			}
		}
	}

	return false
}
