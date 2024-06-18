package leetcode_hot_100

// TODO three - 路径总和 III
func pathSum(root *TreeNode, targetSum int) int {
	prefixMap := map[int]int{0: 1}
	return dfsII(root, targetSum, 0, prefixMap)
}

func dfsII(root *TreeNode, targetSum, curSum int, prefixMap map[int]int) int {
	if root == nil {
		return 0
	}
	curSum = curSum + root.Val
	res := prefixMap[curSum-targetSum]
	prefixMap[curSum]++
	res += dfsII(root.Left, targetSum, curSum, prefixMap) + dfsII(root.Right, targetSum, curSum, prefixMap)
	prefixMap[curSum]--
	return res
}
