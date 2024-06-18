package to_offer

var numStr2NumInt = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var symbolMap = map[string]int{
	" ": 1,
	"+": 1,
	"-": -1,
}

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}

	var (
		start  bool
		result int
		symbol int
	)

	for i := 0; i < len(str); i++ {
		s := string(str[i])

		num, ok := numStr2NumInt[s]
		if !ok && start {
			break
		}

		if !ok {
			symbol, ok = symbolMap[s]
			if !ok {
				break
			}
			if s != " " {
				start = true
			}

			continue
		}

		if result < 2147483648 {
			result = result*10 + num
		}

		start = true
	}

	return getResult(result, symbol)
}

func getResult(num, symbol int) int {
	if symbol == 1 && num >= 2147483647 {
		return 2147483647
	}
	if symbol == -1 && num > 2147483648 {
		return -2147483648
	}
	return num * symbol
}
