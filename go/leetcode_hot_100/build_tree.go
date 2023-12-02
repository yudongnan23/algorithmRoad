package leetcode_hot_100

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := find(inorder, preorder[0])

	root := &TreeNode{
		Val: preorder[0],
	}

	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
}

func find(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		if k == nums[i] {
			return i
		}
	}
	return 0
}
