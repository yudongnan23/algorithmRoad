package to_offer

import "sort"

func checkDynasty(places []int) bool {
	sort.Slice(places, func(i, j int) bool {
		return places[i] < places[j]
	})

	zeroCount := 0
	index := 0

	for index < len(places) {
		if places[index] == 0 {
			zeroCount++
			index++
			continue
		}

		if index == 0 ||
			places[index-1] == 0 ||
			places[index]-places[index-1] == 1 {
			index++
			continue
		}

		if places[index]-places[index-1] == 0 {
			return false
		}

		zeroCount = zeroCount - (places[index] - places[index-1]) + 1
		if zeroCount < 0 {
			return false
		}
		index++
	}

	return true
}
