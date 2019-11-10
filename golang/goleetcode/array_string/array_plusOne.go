package main

/*
	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

	最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

	你可以假设除了整数 0 之外，这个整数不会以零开头。
 */

import "fmt"

func main() {
	digits := []int{9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	if len(digits) == 0{
		return digits
	}
	for i := len(digits) - 1; i >= 0; i--{
		if digits[i] == 9{
			digits[i] = 0
			if i == 0{
				result := []int{}
				result = append([]int{1},digits...)
				return result
			}
		}else{
			digits[i] = digits[i] + 1
			break
		}
	}
	return digits
}
