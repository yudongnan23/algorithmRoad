package leetcode_hot_100

type NodeWithIndex struct {
	Index int
	Val   int
}

// TODO again
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	if n == 0 {
		return res
	}
	stack := make([]NodeWithIndex, 0)
	res[n-1] = 0
	stack = append(stack, NodeWithIndex{
		Index: n - 1,
		Val:   temperatures[n-1],
	})

	for i := n - 2; i >= 0; i-- {
		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j].Val > temperatures[i] {
				res[i] = stack[j].Index - i
				stack = stack[:j+1]
				break
			}
		}

		if res[i] == 0 {
			stack = stack[:0]
		}

		stack = append(stack, NodeWithIndex{
			Index: i,
			Val:   temperatures[i],
		})
	}

	return res
}
