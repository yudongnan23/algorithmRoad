package to_offer

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	nums := []int{5, 3, 4}
	k := 1

	fmt.Println(maxSlidingWindow(nums, k))
}
