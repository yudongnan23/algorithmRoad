package leetcode_hot_100

// TODO again
func isValidBST(root *TreeNode) bool {
	isBst, _, _, _ := bst(root)
	return isBst
}

func bst(root *TreeNode) (bool, bool, int, int) {
	if root == nil {
		return true, true, 0, 0
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false, false, 0, 0
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return false, false, 0, 0
	}

	leftBst, leftEnd, leftMaxVal, leftMinVal := bst(root.Left)
	rightBst, rightEnd, rightMaxVal, rightMinVal := bst(root.Right)

	if !leftBst || !rightBst {
		return false, false, 0, 0
	}

	if !leftEnd && leftMaxVal >= root.Val {
		return false, false, 0, 0
	}

	if !rightEnd && rightMinVal <= root.Val {
		return false, false, 0, 0
	}

	if leftEnd {
		leftMinVal = root.Val
	}

	if rightEnd {
		rightMaxVal = root.Val
	}

	return true, false, rightMaxVal, leftMinVal
}
