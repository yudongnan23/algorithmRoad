package to_offer

func wordPuzzle(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	checked := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		checked[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backTracking(board, checked, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func backTracking(board [][]byte, checked [][]bool, word string, row, col int, index int) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
		return false
	}

	if checked[row][col] {
		return false
	}

	if board[row][col] != word[index] {
		return false
	}

	if len(word)-1 == index {
		return true
	}

	checked[row][col] = true

	res := backTracking(board, checked, word, row-1, col, index+1) ||
		backTracking(board, checked, word, row+1, col, index+1) ||
		backTracking(board, checked, word, row, col-1, index+1) ||
		backTracking(board, checked, word, row, col+1, index+1)

	checked[row][col] = false

	return res
}
