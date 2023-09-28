package to_offer

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	maxNode, minNode := sortNodeII(p, q)

	return backTracing(root, maxNode, minNode)
}

func backTracing(root, maxNode, minNode *TreeNode) *TreeNode {
	if root.Val < maxNode.Val && root.Val > minNode.Val {
		return root
	}

	if root.Val == minNode.Val {
		return minNode
	}

	if root.Val == maxNode.Val {
		return maxNode
	}

	if root.Val < minNode.Val {
		return lowestCommonAncestor(root.Right, minNode, maxNode)
	}

	return lowestCommonAncestor(root.Left, minNode, maxNode)
}

func sortNodeII(p, q *TreeNode) (*TreeNode, *TreeNode) {
	if p.Val > q.Val {
		return p, q
	}
	return q, p
}
