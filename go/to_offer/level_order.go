package to_offer

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		res = append(res, curNode.Val)
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}

		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}

	return res
}
