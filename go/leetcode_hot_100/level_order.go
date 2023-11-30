package leetcode_hot_100

type NodeWithDepth struct {
	Node  *TreeNode
	Depth int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]NodeWithDepth, 1)
	queue[0] = NodeWithDepth{
		Node:  root,
		Depth: 1,
	}

	res := make([][]int, 0)

	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if len(res) < curNode.Depth {
			res = append(res, []int{})
		}

		res[curNode.Depth-1] = append(res[curNode.Depth-1], curNode.Node.Val)

		if curNode.Node.Left != nil {
			queue = append(queue, NodeWithDepth{
				Node:  curNode.Node.Left,
				Depth: curNode.Depth + 1,
			})
		}

		if curNode.Node.Right != nil {
			queue = append(queue, NodeWithDepth{
				Node:  curNode.Node.Right,
				Depth: curNode.Depth + 1,
			})
		}
	}

	return res
}
