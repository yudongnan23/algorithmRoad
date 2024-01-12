package leetcode_hot_100

func kthSmallest(root *TreeNode, k int) int {
	val := -1
	infixOrder(root, &k, &val)
	return val
}

func infixOrder(root *TreeNode, k, val *int) {
	if root == nil {
		return
	}
	infixOrder(root.Left, k, val)
	*k--
	if *k == 0 && *val == -1 {
		*val = root.Val
		return
	}
	infixOrder(root.Right, k, val)
}
