package leetcode_hot_100

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}

	return symmetric(l.Left, r.Right) && symmetric(l.Right, r.Left)
}
