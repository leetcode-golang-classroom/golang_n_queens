package sol

import "strings"

func solveNQueens(n int) [][]string {
	columns := make(map[int]struct{})
	pos_diagonal := make(map[int]struct{})
	neg_diagonal := make(map[int]struct{})

	board := make([][]string, n)
	for row := range board {
		board[row] = make([]string, n)
		for col := 0; col < n; col++ {
			board[row][col] = "."
		}
	}
	result := [][]string{}
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for idx, rowValues := range board {
				temp[idx] = strings.Join(rowValues, "")
			}
			result = append(result, temp)
			return
		}
		for col := 0; col < n; col++ {
			if _, exists := columns[col]; exists {
				continue
			}
			if _, exists := pos_diagonal[row-col]; exists {
				continue
			}
			if _, exists := neg_diagonal[row+col]; exists {
				continue
			}
			columns[col] = struct{}{}
			pos_diagonal[row-col] = struct{}{}
			neg_diagonal[row+col] = struct{}{}
			board[row][col] = "Q"
			dfs(row + 1)
			delete(columns, col)
			delete(pos_diagonal, row-col)
			delete(neg_diagonal, row+col)
			board[row][col] = "."

		}
	}
	dfs(0)
	return result
}
