package to_offer

func deduceTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	leftLen := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			leftLen = i
			break
		}
	}

	root.Left = deduceTree(preorder[1:leftLen+1], inorder[0:leftLen])
	root.Right = deduceTree(preorder[leftLen+1:], inorder[leftLen+1:])

	return root
}
