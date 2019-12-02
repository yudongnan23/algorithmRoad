/*
	给定一个二进制数组， 计算其中最大连续1的个数
*/

package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for variantPoint := 0; variantPoint < len(nums); {
		for variantPoint < len(nums) && nums[variantPoint] != 1 {
			variantPoint++
		}
		count := 0
		for variantPoint < len(nums) && nums[variantPoint] == 1 {
			count++
			variantPoint++
		}
		if result < count || result == 0 {
			result = count
		}
	}
	return result
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
