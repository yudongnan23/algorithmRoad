package leetcode_hot_100

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	val := 0
	infixOrder(root, &count, &val, k)
	return val
}

func infixOrder(node *TreeNode, count, val *int, k int) {
	if *count == k {
		return
	}
	if node == nil {
		return
	}
	infixOrder(node.Left, count, val, k)
	*count++
	if *count == k {
		*val = node.Val
		return
	}
	infixOrder(node.Right, count, val, k)
}
