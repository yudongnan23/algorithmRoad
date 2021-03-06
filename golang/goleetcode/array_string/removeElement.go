/*
	给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素
*/

package main

import "fmt"

func removeElement(nums []int, val int) int {
	var result int
	if len(nums) == 0 {
		return result
	}
	fontPoint := 0
	backPoint := len(nums) - 1
	count := 0
	for fontPoint < backPoint {
		if nums[fontPoint] == val {
			temp := nums[backPoint]
			nums[backPoint] = nums[fontPoint]
			nums[fontPoint] = temp
			backPoint--
		} else {
			count++
			fontPoint++
		}
	}
	if nums[fontPoint] != val {
		count++
	}
	result = len(nums[0:count])
	return result

}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 2
	fmt.Println(removeElement(nums, val))
}
