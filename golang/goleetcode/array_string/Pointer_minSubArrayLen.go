/*
	给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
*/

package main

import "fmt"

// O(n^2)
//func minSubArrayLen(s int, nums []int) int{
//	if len(nums) == 0{
//		return 0
//	}
//	result := 0
//	for i := 0; i < len(nums); {
//		sum := 0
//		count := 0
//		for j := i; j < len(nums);{
//			count++
//			sum = sum + nums[j]
//			if sum >= s{
//				if result > count || result == 0{
//					result = count
//				}
//				break
//			}
//			j++
//		}
//		i++
//
//	}
//	return result
//}

// O(n)
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := 0
	sumAll := 0
	numsLength := len(nums)
	minLength := numsLength + 1
	for left < numsLength {
		if right < numsLength && sumAll < s {
			sumAll += nums[right]
			right++
		} else {
			sumAll = sumAll - nums[left]
			left++
		}
		if sumAll >= s && minLength > right-left {
			minLength = right - left
		}

	}
	if minLength == numsLength+1 {
		return 0
	}
	return minLength

}
func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}
