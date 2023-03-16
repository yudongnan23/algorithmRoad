package to_offer

type TreeNodeWithDepth struct {
	node  *TreeNode
	depth int16
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNodeWithDepth{
		&TreeNodeWithDepth{
			node:  root,
			depth: 1,
		},
	}

	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if len(res) < curNode.depth {
			res = append(res, []int{})
		}

		res[curNode.depth-1] = append(res[curNode.depth-1], curNode.node.Val)

		if curNode.node.Left != nil {
			queue = append(queue, &TreeNodeWithDepth{
				node:  curNode.node.Left,
				depth: curNode.depth + 1,
			})
		}

		if curNode.node.Right != nil {
			queue = append(queue, &TreeNodeWithDepth{
				node:  curNode.node.Right,
				depth: curNode.depth + 1,
			})
		}
	}

	return res
}
