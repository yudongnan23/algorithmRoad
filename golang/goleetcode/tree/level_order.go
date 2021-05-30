package main

import "fmt"

/**
 * Definition for a binary tree node.
 * 层序遍历二叉树
 * 思路: 使用队列实现广度优先遍历
 */

 type TreeNode struct {
 	Val int
   	Left *TreeNode
	Right *TreeNode
 }

func levelOrder(root *TreeNode) [][]int {
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
			result[cur.height - 1] = append(result[cur.height - 1], cur.treeNode.Val)
		}
	}

	return result
}

func main() {
	//root := TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 9,
	//		Left: nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 15,
	//			Left: nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//			Left: nil,
	//			Right: nil,
	//		},
	//	},
	//}

	s := "{\"content\": \"你好\"}"

	fmt.Println([]byte(s))
}
