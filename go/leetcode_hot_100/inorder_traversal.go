package leetcode_hot_100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftVals := inorderTraversal(root.Left)
	leftVals = append(leftVals, root.Val)
	rightVals := inorderTraversal(root.Right)
	leftVals = append(leftVals, rightVals...)
	return leftVals
}
