package leetcode_hot_100

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)

	return root
}
