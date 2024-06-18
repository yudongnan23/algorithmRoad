package to_offer

func reversePairs(record []int) int {
	_, count := dfsIIIIII(record)
	return count
}

func dfsIIIIII(record []int) ([]int, int) {
	if len(record) < 2 {
		return copySlice(record), 0
	}

	mid := len(record) / 2
	left, count1 := dfsIIIIII(record[:mid])
	right, count2 := dfsIIIIII(record[mid:])

	ln := len(left)
	rn := len(right)

	left = append(left, (1<<63)-1)
	right = append(right, (1<<63)-1)

	count := count1 + count2

	res := make([]int, 0)

	li := 0
	ri := 0
	for li < ln || ri < rn {
		if left[li] <= right[ri] {
			count = count + ri
			res = append(res, left[li])
			li++
		} else {
			res = append(res, right[ri])
			ri++
		}
	}

	return res, count
}
