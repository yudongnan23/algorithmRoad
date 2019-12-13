/*
	给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，
	使得从1 到 n 的 min(ai, bi) 总和最大。
 */

/*
    思路：首先需要对数组进行从小到大进行排序，再依次构造两两数对，才能使得所有数对的最小值之和最大
*/

package main

import (
	"fmt"
	"sort"
)


// 方法一，自定义排序函数与min函数
func arrayPairSum(nums []int) int{
	var result int
	if len(nums) == 0{
		return result
	}
	newNums := mySort(nums)
	for i := 0; i <= len(newNums) - 2; i += 2{
		result = result + newNums[i]
	}
 	return result
}

func mySort(nums []int) []int{
	basePoint := 0
	for ; basePoint < len(nums); basePoint++{
		for variantPoint := basePoint; variantPoint < len(nums); variantPoint++{
			if nums[variantPoint] < nums[basePoint] {
				var temp int
				temp = nums[variantPoint]
				nums[variantPoint] = nums[basePoint]
				nums[basePoint] = temp
			}
		}
	}
	return nums
}

func min(num1, num2 int)int{
	if num1 > num2{
		return num2
	}
	return num1
}

func main() {
	nums := []int{1, 3, 12, 4, 53, 1}
	fmt.Println(arrayPairSum(nums))
}


// 方法二，调用sort包中Ints排序函数对切片进行排序
func arrayPairSum(nums []int) int{
	var result int
	if len(nums) == 0{
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i += 2{
		result = result + nums[i]
	}
	return result
}