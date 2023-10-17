package to_offer

// TODO again
func encryptionCalculate(a int, b int) int {
	var (
		sum   int
		carry int
	)

	for b != 0 {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
	}

	return a
}
