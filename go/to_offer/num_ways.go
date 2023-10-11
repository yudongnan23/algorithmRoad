package to_offer

func trainWays(num int) int {
	jjj := 1
	jj := 1
	j := 1

	for num-2 >= 0 {
		jjj = (jj + j) % mod
		jj, j = jjj, jj
		num--
	}

	return jjj
}
