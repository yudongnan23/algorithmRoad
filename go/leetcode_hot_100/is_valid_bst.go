package leetcode_hot_100

// TODO three - 验证二叉搜索树
const (
	minVal = -(1 << 32)
)

func isValidBST(root *TreeNode) bool {
	_, _, isBST := BST(root)
	return isBST
}

func BST(root *TreeNode) (int, int, bool) {
	if root == nil {
		return minVal, minVal, true
	}

	leftMinValue, leftMaxValue, leftBST := BST(root.Left)
	rightMinValue, rightMaxValue, rightBST := BST(root.Right)

	if !leftBST || !rightBST {
		return 0, 0, false
	}

	if (leftMaxValue != minVal && root.Val <= leftMaxValue) ||
		(rightMinValue != minVal && root.Val >= rightMinValue) {
		return 0, 0, false
	}

	if leftMinValue == minVal {
		leftMinValue = root.Val
	}
	if rightMaxValue == minVal {
		rightMaxValue = root.Val
	}

	return leftMinValue, rightMaxValue, true
}
