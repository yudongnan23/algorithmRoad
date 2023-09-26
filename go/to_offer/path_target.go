package to_offer

func pathTarget(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	backTrackingIII(&res, []int{}, 0, root, target)
	return res
}

func backTrackingIII(res *[][]int, curRes []int, curSum int, curNode *TreeNode, target int) {
	curRes = append(curRes, curNode.Val)

	if curNode.Val+curSum == target && curNode.Left == nil && curNode.Right == nil {
		*res = append(*res, curRes)
		return
	}

	if curNode.Left != nil {
		backTrackingIII(res, copySlice(curRes), curNode.Val+curSum, curNode.Left, target)
	}

	if curNode.Right != nil {
		backTrackingIII(res, copySlice(curRes), curNode.Val+curSum, curNode.Right, target)
	}
}

func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
