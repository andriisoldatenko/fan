package main

func main() {}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)
		col := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if ok := row[board[i][j]]; board[i][j] != '.' && ok {
				return false
			}
			row[board[i][j]] = true

			if ok := col[board[j][i]]; board[j][i] != '.' && ok {
				return false
			}
			col[board[j][i]] = true
		}
	}

	// 3x3
	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			subBoxes := make(map[byte]bool)
			for n := x; n < x+3; n++ {
				for m := y; m < y+3; m++ {
					if ok := subBoxes[board[n][m]]; board[n][m] != '.' && ok {
						return false
					}
					subBoxes[board[n][m]] = true
				}
			}
		}
	}
	return true
}
