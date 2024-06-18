package to_offer

// TODO again
func digitOneInNumber(num int) int {
	base := 1
	count := 0

	for base <= num {
		cur := (num / base) % 10
		b := num % base
		a := num / (base * 10)

		if cur < 1 {
			count = count + a*base
		} else if cur == 1 {
			count = count + a*base + b + 1
		} else if cur > 1 {
			count = count + (a+1)*base
		}

		base = base * 10
	}

	return count
}
