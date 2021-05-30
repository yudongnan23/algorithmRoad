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


func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	type treeNodeWithHeight struct {
		treeNode *TreeNode
		height int
	}

	height := 1
	queen := []treeNodeWithHeight{
		{
			treeNode: root,
			height: height,
		},
	}

	for len(queen) != 0 {
		cur := queen[0]
		queen = queen[1:]

		// 若当前节点左子节点不为空，将其入队
		if cur.treeNode.Left != nil {
			queen = append(queen, treeNodeWithHeight{treeNode: cur.treeNode.Left, height: cur.height + 1})
		}
		// 若当前节点右子节点不为空
		if cur.treeNode.Right != nil {
			queen = append(queen, treeNodeWithHeight{treeNode: cur.treeNode.Right, height: cur.height + 1})
		}

		if len(result) != cur.height {
			result = append(result, []int{cur.treeNode.Val})
		}else {
			switch cur.height % 2 {
			case 1:
				result[cur.height - 1] = append(result[cur.height - 1], []int{cur.treeNode.Val}...)
			case 0:
				result[cur.height - 1] = append([]int{cur.treeNode.Val}, result[cur.height - 1]...)
			}
		}
	}

	return result
}


func main() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right:  &TreeNode{
				Val: 5,
				Left: nil,
				Right: nil,
			},
		},
	}

	result := zigzagLevelOrder(&root)
	fmt.Println(result)
}