package leetcode_hot_100

// TODO three - 单词搜索
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board) == 0 {
		return false
	}

	checked := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		checked[i] = make([]bool, len(board[0]))
	}

	var search func(int, int, string) bool
	search = func(i, j int, curWord string) bool {
		if len(curWord) == 0 {
			return true
		}

		if checked[i][j] || board[i][j] != curWord[0] {
			return false
		}

		if len(curWord) == 1 {
			return true
		}

		checked[i][j] = true

		// 上
		if i > 0 && !checked[i-1][j] && search(i-1, j, curWord[1:]) {
			return true
		}

		// 下
		if i < len(board)-1 && !checked[i+1][j] && search(i+1, j, curWord[1:]) {
			return true
		}

		// 左
		if j > 0 && !checked[i][j-1] && search(i, j-1, curWord[1:]) {
			return true
		}

		// 右
		if j < len(board[i])-1 && !checked[i][j+1] && search(i, j+1, curWord[1:]) {
			return true
		}
		checked[i][j] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if search(i, j, word) {
				return true
			}
		}
	}

	return false
}
