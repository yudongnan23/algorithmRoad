package to_offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	queue := []*TreeNode{A}
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if curNode.Val == B.Val && isSub(curNode, B) {
			return true
		}

		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}

		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}

	return false
}

func isSub(A *TreeNode, B *TreeNode) bool {
	queueB := []*TreeNode{B}
	queueA := []*TreeNode{A}

	for len(queueB) != 0 {
		curNodeB := queueB[0]
		curNodeA := queueA[0]

		queueB = queueB[1:]
		queueA = queueA[1:]

		if curNodeB.Val != curNodeA.Val {
			return false
		}

		if curNodeB.Left != nil {
			if curNodeA.Left == nil {
				return false
			}

			queueB = append(queueB, curNodeB.Left)
			queueA = append(queueA, curNodeA.Left)
		}

		if curNodeB.Right != nil {
			if curNodeA.Right == nil {
				return false
			}

			queueB = append(queueB, curNodeB.Right)
			queueA = append(queueA, curNodeA.Right)
		}
	}

	return true
}
