package leetcode_hot_100

func twoSum(nums []int, target int) []int {
	duplicateMapping := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		j, ok := duplicateMapping[target-nums[i]]
		if ok {
			return []int{j, i}
		}

		duplicateMapping[nums[i]] = i
	}

	return []int{}
}
