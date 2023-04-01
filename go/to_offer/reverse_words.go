package to_offer

const (
	_spaceByte byte = 32
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	var res string

	l := 0
	r := 0
	for ; r < len(s); r++ {
		if s[r] == _spaceByte {
			if l < r {
				res = " " + string(s[l:r]) + res
			}
			l = r + 1
		}
	}

	if l < r {
		res = string(s[l:r]) + res
	}

	if len(res) > 0 && res[0] == _spaceByte {
		res = res[1:]
	}

	return res
}
