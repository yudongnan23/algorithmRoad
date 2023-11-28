package leetcode_hot_100

type TreeNodeWithDepth struct {
	Node  *TreeNode
	Depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]TreeNodeWithDepth, 1)
	queue[0] = TreeNodeWithDepth{
		Node:  root,
		Depth: 1,
	}

	maxDepth := 1

	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if curNode.Depth > maxDepth {
			maxDepth = curNode.Depth
		}

		if curNode.Node.Left != nil {
			queue = append(queue, TreeNodeWithDepth{
				Node:  curNode.Node.Left,
				Depth: curNode.Depth + 1,
			})
		}

		if curNode.Node.Right != nil {
			queue = append(queue, TreeNodeWithDepth{
				Node:  curNode.Node.Right,
				Depth: curNode.Depth + 1,
			})
		}
	}

	return maxDepth
}
