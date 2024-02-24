package leetcode_hot_100

// TODO three - 每日温度
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	if length == 0 {
		return res
	}
	stack := make([][2]int, 0)
	res[length-1] = 0
	stack = append(stack, [2]int{temperatures[length-1], length - 1})
	for i := length - 2; i >= 0; i-- {
		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j][0] > temperatures[i] {
				res[i] = stack[j][1] - i
				stack = stack[:j+1]
				break
			}
		}

		if res[i] == 0 {
			stack = stack[:0]
		}
		stack = append(stack, [2]int{temperatures[i], i})
	}

	return res
}
