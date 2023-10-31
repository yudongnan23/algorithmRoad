package leetcode_hot_100

// TODO again
func trap(height []int) int {
	length := len(height)

	if length == 0 {
		return 0
	}

	leftMax := make([]int, length)
	rightMax := make([]int, length)

	leftMax[0] = height[0]
	for i := 1; i < length; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	rightMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	area := 0

	for i := 0; i < length; i++ {
		area = area + (min(leftMax[i], rightMax[i]) - height[i])
	}

	return area
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
