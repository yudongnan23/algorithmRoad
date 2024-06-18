package to_offer

type TreeNodeWithDepthIII struct {
	Node *TreeNode

	Depth int
}

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []TreeNodeWithDepthIII{
		{
			Node:  root,
			Depth: 1,
		},
	}

	var maxDepth int

	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if maxDepth < curNode.Depth {
			maxDepth = curNode.Depth
		}

		if curNode.Node.Left != nil {
			queue = append(queue, TreeNodeWithDepthIII{Node: curNode.Node.Left, Depth: curNode.Depth + 1})
		}

		if curNode.Node.Right != nil {
			queue = append(queue, TreeNodeWithDepthIII{Node: curNode.Node.Right, Depth: curNode.Depth + 1})
		}
	}

	return maxDepth
}
