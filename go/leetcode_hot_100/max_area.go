package leetcode_hot_100

// TODO again
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		curArea := (right - left) * min(height[left], height[right])
		if curArea > max {
			max = curArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
