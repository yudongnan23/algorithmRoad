package to_offer

func fileCombination(target int) [][]int {
	res := make([][]int, 0)

	for i := 1; i <= target/2; i++ {
		sum := 0
		for j := i; j <= target; j++ {
			sum = sum + j
			if sum > target {
				break
			}

			if sum == target {
				curRes := make([]int, j-i+1)
				for x := i; x <= j; x++ {
					curRes[x-i] = x
				}
				res = append(res, curRes)
				break
			}
		}
	}

	return res
}
