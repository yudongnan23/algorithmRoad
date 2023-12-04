package leetcode_hot_100

// TODO again
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _ := dfsIII(root, p, q)
	return node
}

func dfsIII(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}

	equal := root == p || root == q

	leftNode, leftFound := dfsIII(root.Left, p, q)
	if leftNode != nil {
		return leftNode, true
	}

	if equal && leftFound {
		return root, true
	}

	rightNode, rightFound := dfsIII(root.Right, p, q)
	if rightNode != nil {
		return rightNode, true
	}

	if equal && rightFound {
		return root, true
	}

	if rightFound && leftFound {
		return root, true
	}

	if equal {
		return nil, true
	}

	if rightFound || leftFound {
		return nil, true
	}

	return nil, false
}
