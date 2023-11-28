package leetcode_hot_100

// TODO again
func diameterOfBinaryTree(root *TreeNode) int {
	var maxLength int
	_ = diameter(root, &maxLength)
	if maxLength <= 1 {
		return maxLength
	}
	return maxLength - 1
}

func diameter(root *TreeNode, maxLength *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	maxLeft := diameter(root.Left, maxLength)
	maxRight := diameter(root.Right, maxLength)

	maxLengthWithSelf := maxLeft + maxRight + 1
	if maxLengthWithSelf > *maxLength {
		*maxLength = maxLengthWithSelf
	}

	if maxLeft > maxRight {
		return maxLeft + 1
	}
	return maxRight + 1
}
