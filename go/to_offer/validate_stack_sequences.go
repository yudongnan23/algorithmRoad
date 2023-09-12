package to_offer

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	if len(pushed) == 0 || len(popped) == 0 {
		return false
	}

	mockStack := make([]int, 0)
	popIndex := 0

	for pushIndex := 0; pushIndex < len(pushed); pushIndex++ {
		mockStack = append(mockStack, pushed[pushIndex])

		for {
			if len(mockStack) == 0 {
				break
			}

			if mockStack[len(mockStack)-1] == popped[popIndex] {
				popIndex++
				mockStack = mockStack[:(len(mockStack) - 1)]
				continue
			}

			break
		}
	}

	if len(mockStack) != len(popped[popIndex:]) {
		return false
	}

	j := len(mockStack) - 1

	for i := 0; i < len(mockStack); i++ {
		if mockStack[i] != popped[j] {
			return false
		}

		j--
	}

	return true
}
