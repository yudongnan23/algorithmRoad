/*
	给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数
*/
package main

import "fmt"

//func rotate(nums []int, k int) []int{
//	if len(nums) == 0{
//		return nums
//	}
//	moveStart := len(nums) - 1
//	for count := 0; count < k; count++{
//		for variant := moveStart; variant > 0; variant--{
//			temp := nums[variant-1]
//			nums[variant-1] = nums[variant]
//			nums[variant] = temp
//		}
//	}
//	return nums
//}

func rotate(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	numsLength := len(nums)
	if numsLength == 2 {
		if k%2 == 1 {
			temp := nums[1]
			nums[1] = nums[0]
			nums[0] = temp
		}
		return nums
	}
	reverse, moveNum := numsHandler(numsLength, k)
	if reverse {
		reversed(nums)
	}
	for count := 0; count < moveNum; count++ {
		for variant := numsLength - 1; variant > 0; variant-- {
			temp := nums[variant]
			nums[variant] = nums[variant-1]
			nums[variant-1] = temp
		}
	}
	return nums
}

func reversed(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		temp := nums[right]
		nums[right] = nums[left]
		nums[left] = temp
		left++
		right--
	}
}

func numsHandler(numsLength int, k int) (bool, int) {
	var reverse bool
	multiple := k / (numsLength - 1)
	if multiple%2 == 0 {
		reverse = false
	} else {
		reverse = true
	}
	moveNums := k % numsLength
	return reverse, moveNums

}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	fmt.Println(rotate(nums, k))
}
