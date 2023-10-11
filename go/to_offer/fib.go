package to_offer

const mod = 1000000007

func fib(n int) int {
	if n < 2 {
		return n
	}

	jjj := 0
	j := 0
	jj := 1

	for n-2 >= 0 {
		jjj = (j + jj) % mod
		jj, j = jjj, jj
		n--
	}

	return jjj
}
