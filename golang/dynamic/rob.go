/*
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
    如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额
 */

package main

func rob(nums []int)int{
	length := len(nums)
	if length == 0{
		return 0
	}
	if length == 1{
		return nums[0]
	}
	if length == 2{
		return getMaxValue(nums[0], nums[1])
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = getMaxValue(nums[0], nums[1])
	maxValue := 0
	for i := 2; i < length; i++{
		dp[i] = getMaxValue(dp[i - 1], dp[i - 2] + nums[i])

		maxValue = getMaxValue(maxValue, dp[i])
	}
	return maxValue
}

func getMaxValue(value1 int, value2 int)int{
	if value1 > value2{
		return value1
	}else{
		return value2
	}
}