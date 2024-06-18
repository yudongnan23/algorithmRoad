package to_offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	turnOver(root)
	return root
}

func turnOver(root *TreeNode) {
	if root == nil {
		return
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	turnOver(root.Left)
	turnOver(root.Right)
}
