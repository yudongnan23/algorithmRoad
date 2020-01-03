/*
	给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 */


/*
    思路： 本题采用广度优先搜索方法，定义一个具有队列性质的列表，该队列以列表为元素，列表包含两个元素，第一个元素为搜索过的数的平方和，第二个元素为搜索过的完全平方数的个数，初始化为[[0, 0]]。
            弹出队列队首元素，获得当前平方和以及得到该平方和的完全平方数的个数，对平方小于或等于目标数的数进行遍历（搜索）,将该数的平方加到原有的平方和得到新的平方和，
            再将完全平方数的个数加1得到新的完全平方数的个数，重复上述步骤，直到平方和等于目标数，跳出循环，返回和个数；需要注意的是，为了避免平方和重复进入队列，
            需要建立一个记录已搜索得到的平方和的字典，将不存在该字典以及平方和小于目标数的平方和以及得到平方和的完全平方和数的个数添加进入队列，再将平方和添加进入记录已遍历的平方和的字典
*/

package main

import (
	"fmt"
)

func numSquares(n int) int {
	var step int
	var breakout bool
	// 声明并初始化记录已经搜索得到的平方和的字典
	dict := make(map[int]bool)
	queue := [][]int{[]int{0, 0}}

	if n < 1{
		return 0
	}
	for len(queue) != 0{
	    // 弹出队首元素操作，得到平方和以及得到该平方和的完全平方数的个数
		curSum := queue[0][0]
		curStep := queue[0][1]
		queue = queue[1:]

        // 将平方小于目标数的数进行遍历搜索
		for value := 1; value*value <= n; value++{
		    // 得到新的平方和
			newSum := curSum + value*value
			// 得到新的完全平方数个数
			newStep := curStep + 1

			// 若新的平方和等于目标数，直接退出循环，内层循环以及外层循环，返回完全平方数的个数
			if newSum == n{
				step = newStep
				breakout = true
				break
			}

			// 若新的平方和不存在于字典并且新的平方和小于目标数，进行新元素入队操作，再将新的平方和元素添加进入字典
			if _, exist := dict[newSum]; !exist  && newSum < n{
				queue = append(queue, []int{newSum, newStep})
				dict[newSum] = true
			}
		}

		if breakout{
			break
		}

	}
	return step
}

func main() {
    // 测试用例
	sum := 13
	fmt.Println(numSquares(sum))
}