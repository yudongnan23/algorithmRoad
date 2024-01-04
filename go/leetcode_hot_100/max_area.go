package leetcode_hot_100

// TODO again
func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		minHeight := height[left]
		length := right - left
		if minHeight > height[right] {
			minHeight = height[right]
			right--
		} else {
			left++
		}

		curArea := length * minHeight
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}
