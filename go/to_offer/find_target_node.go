package to_offer

func findTargetNode(r *TreeNode, k int) int {
	if r == nil {
		return 0
	}

	if k == 0 {
		return 0
	}

	res := make([]int, 0)
	dfs(r, &res)

	return res[len(res)-k]
}

func dfs(r *TreeNode, res *[]int) {
	if r == nil {
		return
	}

	if r.Left != nil {
		dfs(r.Left, res)
	}

	*res = append(*res, r.Val)

	if r.Right != nil {
		dfs(r.Right, res)
	}

	return
}
