package to_offer

func verifyTreeOrder(postorder []int) bool {
	return dfsIIIII(postorder, 0, len(postorder)-1)
}

func dfsIIIII(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	i := start

	for i < end {
		if postorder[i] > postorder[end] {
			break
		}
		i++
	}

	j := i

	for j < end {
		if postorder[j] < postorder[end] {
			return false
		}
		j++
	}

	return dfsIIIII(postorder, start, i-1) && dfsIIIII(postorder, i, end-1)
}
