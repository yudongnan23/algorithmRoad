/*
	给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
	返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
 */


package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	var ways int
	if len(nums) == 0 {
		return ways
	}

	multiple := []int{1, -1}
	maxNumIndex := len(nums) - 1
	stack := [][]int{[]int{nums[0], 0}, []int{-nums[0], 0}}

	for len(stack) != 0{
		curSum := stack[len(stack) - 1][0]
		curNumIndex := stack[len(stack) - 1][1]
		stack = stack[0:len(stack)-1]

		if curSum == S && curNumIndex == maxNumIndex{
			ways ++
			continue
		}

		nextNumIndex := curNumIndex + 1

		if nextNumIndex > maxNumIndex {
			continue
		}
		for _, value := range multiple{
			nextSum := curSum + nums[nextNumIndex]*value
			if nextNumIndex == maxNumIndex{
				if nextSum == S{
					ways ++
				}
				continue
			}

			stack = append(stack, []int{nextSum, nextNumIndex})
		}
	}

	return ways
}

func FindTargetSumWays(nums []int, S int) int{
	var ways int
	if len(nums) == 0{
		return ways
	}

	curSum := 0
	dfs(nums, curSum, S, &ways)

	return ways
}

func dfs(nums []int, curSum int, S int, ways *int){
	numsLength := len(nums)

	if numsLength == 0{
		if curSum == S {
			*ways ++
		}
		return
	}

	multiple := []int{1, -1}
	for _, value := range multiple{
		newSum := curSum + nums[numsLength - 1] * value
		newNums := nums[0:numsLength-1]
		dfs(newNums, newSum, S, ways)
	}
}

func main() {
	nums := []int{6,33,25,41,15,46,36,11,29,18,17,26,28,11,39,4,19,13,40,8}
	S := 5
	fmt.Println(FindTargetSumWays(nums, S))
}