package leetcode_hot_100

// TODO three - 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	targetNode, _ := dfsI(root, p, q)
	return targetNode
}

func dfsI(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	rootExist := (root == p) || (root == q)
	targetNode, leftExist := dfsI(root.Left, p, q)
	if targetNode != nil {
		return targetNode, true
	}
	if leftExist && rootExist {
		return root, true
	}
	targetNode, rightExist := dfsI(root.Right, p, q)
	if targetNode != nil {
		return targetNode, true
	}
	if rootExist && rightExist {
		return root, true
	}
	if leftExist && rightExist {
		return root, true
	}

	return nil, rootExist || leftExist || rightExist
}
