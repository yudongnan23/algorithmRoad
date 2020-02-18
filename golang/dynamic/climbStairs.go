/*
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
    每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 */

package main

func climbStairs(n int)int{
	if n == 0{
		return 1
	}
	// 开一个大小为n + 1的数组
	dp := make([]int, n + 1)
	//  由题意可知，数组的第0个元素和第1个元素都为1
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n + 1; i++{
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}