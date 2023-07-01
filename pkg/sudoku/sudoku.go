package sudoku

func SolveSudoku(board *[9][9]int) {
	recursion(board, 0, 0)
}

func recursion(board *[9][9]int, i, j int) bool {
	if i == 9 {
		return true
	}
	if j >= 9 {
		return recursion(board, i+1, 0)
	}
	if board[i][j] != 0 {
		return recursion(board, i, j+1)
	}
	for c := 1; c <= 9; c++ {
		if !isValid(*board, i, j, c) {
			continue
		}
		board[i][j] = c
		if recursion(board, i, j+1) {
			return true
		}
		board[i][j] = 0
	}
	return false
}

func IsValidSudoku(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				continue
			}
			if !isValid(board, i, j, board[i][j]) {
				return false
			}
		}
	}
	return true
}

func isValid(board [9][9]int, i, j int, val int) bool {
	for x := 0; x < 9; x++ {
		if x == i {
			continue
		}
		if board[x][j] == val {
			return false
		}
	}
	for y := 0; y < 9; y++ {
		if y == j {
			continue
		}
		if board[i][y] == val {
			return false
		}
	}
	row, col := i-i%3, j-j%3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x+row == i && y+col == j {
				continue
			}
			if board[x+row][y+col] == val {
				return false
			}
		}
	}
	return true
}
