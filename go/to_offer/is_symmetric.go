package to_offer

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetry(root.Left, root.Right)
}

func symmetry(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return symmetry(a.Left, b.Right) && symmetry(a.Right, b.Left)
}
