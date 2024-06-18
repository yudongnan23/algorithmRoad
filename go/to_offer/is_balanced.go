package to_offer

func isBalanced(root *TreeNode) bool {
	balanced, _ := dfsII(root)
	return balanced
}

func dfsII(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftDepth := dfsII(root.Left)
	rightBalanced, rightDepth := dfsII(root.Right)

	if !leftBalanced || !rightBalanced {
		return false, 0
	}

	if leftDepth-rightDepth > 1 || leftDepth-rightDepth < -1 {
		return false, 0
	}

	return true, maxDepth(leftDepth, rightDepth) + 1
}

func maxDepth(leftDepth, rightDepth int) int {
	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}
