package to_offer

import (
	"fmt"
)

var numMapping = map[string]bool{
	"0":  true,
	"2":  true,
	"1":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}

	countLeft := 1
	countRight := 1
	numStr := fmt.Sprintf("%d", num)
	for i := 1; i < len(numStr); i++ {
		if numMapping[string(numStr[i-1:i+1])] {
			countRight, countLeft = countRight+countLeft, countRight
		} else {
			countLeft = countRight
		}
	}
	return countRight
}
