package leetcode_hot_100

// TODO again
func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	right := root.Right
	leftLastNode := dfs(root.Left)
	if leftLastNode != nil {
		root.Right = root.Left
		leftLastNode.Right = right
	}
	root.Left = nil
	rightLastNode := dfs(right)

	if rightLastNode == nil {
		rightLastNode = leftLastNode
	}

	if rightLastNode == nil {
		rightLastNode = root
	}

	return rightLastNode
}
