package leetcode_hot_100

// TODO three - 前 K 个高频元素
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	length := len(countMap)
	newNums := make([][2]int, 0, length)
	for num, count := range countMap {
		newNums = append(newNums, [2]int{num, count})
	}
	for i := length / 2; i >= 0; i-- {
		heapII(newNums, i, length)
	}

	sortCount := 0
	for i := len(newNums) - 1; i > 0 && sortCount < k; i-- {
		swapII(newNums, 0, i)
		length--
		heapII(newNums, 0, length)
		sortCount++
	}

	res := make([]int, 0, k)
	for i := len(newNums) - 1; len(res) < k; i-- {
		res = append(res, newNums[i][0])
	}
	return res
}

func heapII(nums [][2]int, index, length int) {
	left := 2*index + 1
	right := 2*index + 2

	largest := index

	if left < length && nums[left][1] > nums[largest][1] {
		largest = left
	}

	if right < length && nums[right][1] > nums[largest][1] {
		largest = right
	}

	if largest != index {
		swapII(nums, largest, index)
		heapII(nums, largest, length)
	}
}

func swapII(nums [][2]int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
