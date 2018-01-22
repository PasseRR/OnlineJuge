package page2

// 常量
const (
	POINT = '.'
	QUEEN = 'Q'
)

// 结果
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := [][]byte{}
	for i := 0; i < n; i++ {
		row := []byte{}
		for j := 0; j < n; j++ {
			row = append(row, POINT)
		}
		queens = append(queens, row)
	}
	backtrackingFillQueens(queens, 0)
	return solutions
}

func backtrackingFillQueens(queens [][]byte, row int) bool {
	length := len(queens)
	for i := 0; i < length; i++ {
		if isLegal(queens, row, i) {
			queens[row][i] = QUEEN
			// 到最后一行 说明皇后已经填充完毕
			if row == length-1 {
				// change [][]byte to []string
				solution := []string{}
				for _, value := range queens {
					solution = append(solution, string(value))
				}
				solutions = append(solutions, solution)
			} else {
				backtrackingFillQueens(queens, row+1)
			}
			// 回溯
			queens[row][i] = POINT
		}
	}

	return false
}

// x为一维索引 y为二维索引
// 新增一个queen 判断是否是合法的
func isLegal(queens [][]byte, x, y int) bool {
	length := len(queens)
	for i := 0; i < length; i++ {
		// 水平方向没有queen
		// 垂直方向没有queen
		if queens[x][i] == QUEEN || queens[i][y] == QUEEN {
			return false
		}
	}

	// 两条对角线上没有queen
	// 以插入索引为中心的四段对角线
	for i, j := x, y; i < length && j < length; i, j = i+1, j+1 {
		if queens[i][j] == QUEEN {
			return false
		}
	}
	for i, j := x, y; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queens[i][j] == QUEEN {
			return false
		}
	}
	for i, j := x, y; i >= 0 && j < length; i, j = i-1, j+1 {
		if queens[i][j] == QUEEN {
			return false
		}
	}
	for i, j := x, y; i < length && j >= 0; i, j = i+1, j-1 {
		if queens[i][j] == QUEEN {
			return false
		}
	}
	return true
}
