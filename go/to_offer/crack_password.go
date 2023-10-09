package to_offer

import (
	"fmt"
	"sort"
)

func crackPassword(password []int) string {
	if len(password) == 0 {
		return ""
	}

	sort.Slice(password, func(i, j int) bool {
		return less(password[i], password[j])
	})

	var res string
	for i := 0; i < len(password); i++ {
		res = res + fmt.Sprintf("%d", password[i])
	}

	return res
}

func less(i, j int) bool {
	si := fmt.Sprintf("%d", i)
	sj := fmt.Sprintf("%d", j)

	if si+sj > sj+si {
		return false
	}

	return true
}
