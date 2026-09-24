func isValidSudoku(board [][]byte) bool {

	colMap := make([][]bool, 9)
	rowMap := make([][]bool, 9)
	sqtMap := make([][]bool, 9)
	for i := range 9 {
		colMap[i] = make([]bool, 9)
		rowMap[i] = make([]bool, 9)
		sqtMap[i] = make([]bool, 9)
	}

	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}

			boardInt := int(board[i][j] - '1')
			if colMap[j][boardInt] {
				return false
			}
			colMap[j][boardInt] = true

			if rowMap[i][boardInt] {
				return false
			}
			rowMap[i][boardInt] = true
			
			s := sqtValue(i, j)
			if sqtMap[s][boardInt] {
				return false
			}
			sqtMap[s][boardInt] = true
		}
	}
	return true
}

func sqtValue(x, y int) int {
	if x < 3 && y < 3 {
		return 0
	} else if x < 6 && y < 3 {
		return 1
	} else if y < 3 {
		return 2
	} else if x < 3 && y < 6 {
		return 3
	} else if x < 6 && y < 6 {
		return 4
	} else if y < 6 {
		return 5
	} else if x < 3 {
		return 6
	} else if x < 6 {
		return 7
	}
	return 8
}
