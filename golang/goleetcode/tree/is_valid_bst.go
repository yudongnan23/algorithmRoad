package main

import "html/template"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftResult := true
	rightResult := true
	if root.Left != nil {
		leftResult = isValidBST(root.Left) && root.Val > root.Left.Val
	}
	if root.Right != nil {
		rightResult = isValidBST(root.Right) && root.Val < root.Right.Val
	}

	return leftResult && rightResult
}