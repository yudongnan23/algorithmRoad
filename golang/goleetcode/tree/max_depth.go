package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var depth int
	var leftDepth int
	var rightDepth int
	if root == nil {
		return 0
	}

	leftDepth = maxDepth(root.Left) + 1
	rightDepth = maxDepth(root.Right) + 1

	if leftDepth > rightDepth{
		depth = leftDepth
	}else {
		depth = rightDepth
	}

	return depth
}

func main() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 7,
				Left: nil,
				Right: nil,
			},
		},
	}

	result := maxDepth(&root)
	fmt.Println(result)
}
