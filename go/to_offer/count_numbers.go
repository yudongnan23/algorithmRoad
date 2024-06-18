package to_offer

func countNumbers(cnt int) []int {
	maxNumer := getMaxNumber(cnt)
	res := make([]int, maxNumer-1)

	for i := 1; i < maxNumer; i++ {
		res[i-1] = i
	}

	return res
}

func getMaxNumber(cnt int) int {
	if cnt == 0 {
		return 1
	}
	temp := getMaxNumber(cnt / 2)
	if cnt%2 == 1 {
		return temp * temp * 10
	}
	return temp * temp
}
