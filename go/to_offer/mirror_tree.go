package to_offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		curNode.Left, curNode.Right = curNode.Right, curNode.Left

		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}

		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}

	return root
}
