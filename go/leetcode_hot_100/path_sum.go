package leetcode_hot_100

// TODO again
func pathSum(root *TreeNode, targetSum int) int {
	var count int
	trapII(root, targetSum, &count)
	return count
}

func trapII(root *TreeNode, target int, count *int) {
	if root == nil {
		return
	}

	dfsII(root, target, 0, count)
	trapII(root.Left, target, count)
	trapII(root.Right, target, count)
}

func dfsII(root *TreeNode, target int, curSum int, count *int) {
	if root == nil {
		return
	}

	if curSum+root.Val == target {
		*count++
	}

	dfsII(root.Left, target, curSum+root.Val, count)
	dfsII(root.Right, target, curSum+root.Val, count)
}
