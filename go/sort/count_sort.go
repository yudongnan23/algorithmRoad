package sort

func countSort(nums []int) {
	length := len(nums)
	minNum, maxNum := numMaxAndMin(nums)
	countArrayLength := maxNum - minNum + 1
	countArray := make([]int, countArrayLength)
	for i := 0; i < length; i++ {
		countArray[nums[i]-minNum]++
	}

	index := 0
	for i := 0; i < countArrayLength; i++ {
		if countArray[i] == 0 {
			continue
		}
		for j := countArray[i]; j > 0; j-- {
			nums[index] = i + minNum
			index++
		}
	}
}

func numMaxAndMin(nums []int) (int, int) {
	length := len(nums)
	if length == 0 {
		return 0, 0
	}
	minNum := nums[0]
	maxNum := nums[0]
	for i := 0; i < length; i++ {
		if minNum > nums[i] {
			minNum = nums[i]
		}

		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}
	return minNum, maxNum
}
