package leetcode_hot_100

const (
	minNum = -(2 << 32)
)

// TODO three
func maxPathSum(root *TreeNode) int {
	maxLength := -(2 << 32)
	maxLen := dfsIIII(root, &maxLength)
	if maxLen > maxLength {
		maxLength = maxLen
	}
	return maxLength
}

func dfsIIII(root *TreeNode, maxLength *int) int {
	if root == nil {
		return minNum
	}

	leftMaxLength := dfsIIII(root.Left, maxLength)
	rightMaxLength := dfsIIII(root.Right, maxLength)

	curMaxLength := maxInTwo(maxInTwo(leftMaxLength+rightMaxLength+root.Val, leftMaxLength), rightMaxLength)
	if curMaxLength > *maxLength {
		*maxLength = curMaxLength
	}

	return maxInTwo(maxInTwo(rightMaxLength+root.Val, leftMaxLength+root.Val), root.Val)
}

func maxInTwo(x, y int) int {
	if x > y {
		return x
	}
	return y
}
