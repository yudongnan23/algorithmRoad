package leetcode_hot_100

// TODO three
func largestRectangleArea(heights []int) int {
	n := len(heights)
	area := 0
	stack := make([]int, 0, n+1)
	stack = append(stack, -1)
	heights = append(heights, 0)
	n++

	for i := 0; i < n; i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			newArea := (i - left - 1) * heights[right]
			if newArea > area {
				area = newArea
			}
		}
		stack = append(stack, i)
	}
	return area
}
