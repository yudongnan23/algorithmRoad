/*
	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
	最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
	你可以假设除了整数 0 之外，这个整数不会以零开头。
 */

/*
    思路：从切片的最后一个元素开始遍历，当遇到9时，需要对当前元素进行改1操作，进行下一个元素的遍历；当遇到的不是9时，只需要将当前遍历元素加1，然后结束遍历，返回切片即可
*/

package main

import "fmt"

func main() {
    // 测试用例
	digits := []int{9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	if len(digits) == 0{
		return digits
	}
	for i := len(digits) - 1; i >= 0; i--{
	    // 当遇到9时
		if digits[i] == 9{
		    // 当前元素直接改1
			digits[i] = 0
			// 如果数组元素最大索引为0，即切片只存在一个元素9，将1添加至切片头部，直接返回切片即可
			if i == 0{
				result := []int{}
				result = append([]int{1},digits...)
				return result
			}
		}else{
		    // 遍历得到的元素不是9时，将当前元素加1直接结束遍历
			digits[i] = digits[i] + 1
			break
		}
	}
	return digits
}
