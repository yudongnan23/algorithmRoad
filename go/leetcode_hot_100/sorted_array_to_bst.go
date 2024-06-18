package leetcode_hot_100

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	midIndex := len(nums) / 2
	node := &TreeNode{
		Val:   nums[midIndex],
		Left:  sortedArrayToBST(nums[:midIndex]),
		Right: sortedArrayToBST(nums[midIndex+1:]),
	}

	return node
}
