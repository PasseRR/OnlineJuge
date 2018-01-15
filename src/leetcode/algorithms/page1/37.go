package page1

func solveSudoku(board [][]byte) {
	fill(board)
}

func isValidByte(row, col int, b byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == b || // row check
			board[i][col] == b || // column check
			board[row/3*3+i/3][col/3*3+i%3] == b { // sub-boxes check
			return false
		}
	}

	return true
}

func fill(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for b := byte('1'); b <= byte('9'); b++ {
					if isValidByte(i, j, b, board) {
						board[i][j] = b
						if fill(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}

				return false
			}
		}
	}

	return true
}
