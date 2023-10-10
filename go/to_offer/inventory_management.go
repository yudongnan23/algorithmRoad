package to_offer

import "sort"

func inventoryManagement(stock []int, cnt int) []int {
	sort.Slice(stock, func(i, j int) bool {
		return stock[i] < stock[j]
	})

	return stock[:cnt]
}
