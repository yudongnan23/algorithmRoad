/*
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 */

package main

func maxSubArray(nums []int)int{
	length := len(nums)
	if length == 0{
		return 0
	}
	// 开一个大小为length的动态数组，且其元素都初始化为0
	dp := make([]int, length)
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i< length; i++{
		add := dp[i - 1] + nums[i]
		notAdd := nums[i]
		if add > notAdd{
			dp[i] = add
		}else{
			dp[i] = notAdd
		}
		if maxSum < dp[i]{
			maxSum = dp[i]
		}
	}
	return maxSum
}
