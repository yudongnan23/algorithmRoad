package to_offer

func statisticalResult(arrayA []int) []int {
	if len(arrayA) == 0 {
		return []int{}
	}

	res := make([]int, len(arrayA))

	sum := 1

	zeroIndex := make(map[int]bool, 0)

	for i := 0; i < len(arrayA); i++ {
		if arrayA[i] != 0 {
			sum = arrayA[i] * sum
		} else {
			zeroIndex[i] = true
		}
	}

	for i := 0; i < len(arrayA); i++ {
		if arrayA[i] != 0 {
			if len(zeroIndex) > 0 {
				res[i] = 0
			} else {
				res[i] = sum / arrayA[i]
			}

			continue
		}

		if len(zeroIndex) == 1 {
			res[i] = sum
		} else {
			res[i] = 0
		}
	}

	return res
}
