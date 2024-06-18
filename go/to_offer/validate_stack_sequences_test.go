package to_offer

import (
	"fmt"
	"testing"
)

func Test_validateStackSequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	poped := []int{4, 5, 3, 2, 1}

	fmt.Println(validateStackSequences(pushed, poped))
}
