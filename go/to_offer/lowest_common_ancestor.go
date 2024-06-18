package to_offer

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	node, _ := dfsIII(root, p, q)
	if node == nil {
		return root
	}
	return node
}

func dfsIII(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root == p {
		return p, true
	}

	if root == q {
		return q, true
	}

	leftNode, leftExist := dfsIII(root.Left, p, q)
	rightNode, rightExist := dfsIII(root.Right, p, q)

	if leftExist && rightExist {
		return root, true
	}

	if leftExist {
		return leftNode, true
	}

	if rightExist {
		return rightNode, true
	}

	return nil, false
}
