package to_offer

func reverseLeftWords(s string, n int) string {
	return string(s[n:]) + string(s[:n])
}
