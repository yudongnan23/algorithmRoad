/*
	给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
	函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
	说明:
		返回的下标值（index1 和 index2）不是从零开始的。
		你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
*/
package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	fontPoint := 0
	backPoint := len(numbers) - 1
	for fontPoint < backPoint {
		if numbers[fontPoint]+numbers[backPoint] == target {
			result := []int{fontPoint + 1, backPoint + 1}
			return result
		} else if numbers[fontPoint]+numbers[backPoint] > target {
			backPoint--
			continue
		} else if numbers[fontPoint]+numbers[backPoint] < target {
			fontPoint++
		}
	}
	return []int{}

}

func main() {
	testSum := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(testSum, target))
}
