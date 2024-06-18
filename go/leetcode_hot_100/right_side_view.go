package leetcode_hot_100

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []NodeWithDepth{{Node: root, Depth: 0}}
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode.Depth == len(res) {
			res = append(res, 0)
		}
		res[curNode.Depth] = curNode.Node.Val

		if curNode.Node.Left != nil {
			queue = append(queue, NodeWithDepth{Node: curNode.Node.Left, Depth: curNode.Depth + 1})
		}

		if curNode.Node.Right != nil {
			queue = append(queue, NodeWithDepth{Node: curNode.Node.Right, Depth: curNode.Depth + 1})
		}
	}

	return res
}
