package page1

func initOccurCount() map[byte]bool {
	return map[byte]bool{
		'1': false,
		'2': false,
		'3': false,
		'4': false,
		'5': false,
		'6': false,
		'7': false,
		'8': false,
		'9': false,
	}
}

func isValidSudoku(board [][]byte) bool {
	size := len(board)
	// Each row must have the numbers 1-9 occuring just once.
	// Each column must have the numbers 1-9 occuring just once.
	for i := 0; i < size; i++ {
		rowCount, colCount := initOccurCount(), initOccurCount()
		for j := 0; j < size; j++ {
			rowValue, colValue := board[i][j], board[j][i]
			if rowValue != '.' {
				if rowCount[rowValue] {
					return false
				} else {
					rowCount[rowValue] = true
				}
			}

			if colValue != '.' {
				if colCount[colValue] {
					return false
				} else {
					colCount[colValue] = true
				}
			}
		}
	}

	// And the numbers 1-9 must occur just once in each of the 9 sub-boxes of the grid.
	for i := 0; i < size; i += 3 {
		for j := 0; j < size; j += 3 {
			count := initOccurCount()
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					value := board[k][l]
					if value != '.' {
						if count[value] {
							return false
						} else {
							count[value] = true
						}
					}
				}
			}
		}
	}

	return true
}
