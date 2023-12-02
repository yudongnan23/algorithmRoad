package leetcode_hot_100

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	levelRange := make([][]int, 0)
	queue := make([]NodeWithDepth, 1)
	queue[0] = NodeWithDepth{
		Node:  root,
		Depth: 1,
	}

	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if len(levelRange) < curNode.Depth {
			levelRange = append(levelRange, []int{})
		}

		levelRange[curNode.Depth-1] = append(levelRange[curNode.Depth-1], curNode.Node.Val)

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

	res := make([]int, len(levelRange))
	for i := 0; i < len(levelRange); i++ {
		res[i] = levelRange[i][len(levelRange[i])-1]
	}

	return res
}
