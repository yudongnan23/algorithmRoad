package to_offer

func hammingWeight(n uint32) int {
	var count int
	for n > 0 {
		if n%2 == 1 {
			count++
		}

		n = n >> 1
	}

	return count
}
