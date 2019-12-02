/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
*/

package main

import "fmt"

func moveZeros(nums []int) {
	if len(nums) == 0 {

	} else {
		leftPoint := -1
		rightPoint := 0
		for rightPoint < len(nums) && leftPoint < len(nums)-1 {
			if nums[leftPoint+1] == 0 {
				for ; rightPoint < len(nums); rightPoint++ {
					if nums[rightPoint] != 0 {
						temp := nums[rightPoint]
						nums[rightPoint] = nums[leftPoint+1]
						nums[leftPoint+1] = temp
						leftPoint++
					}
				}
			} else {
				leftPoint++
				rightPoint++
			}
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{1, 1, 0, 3, 12}
	moveZeros(nums)
}
