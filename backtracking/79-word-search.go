package backtracking

var moves = [][]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

// https://leetcode.com/problems/word-search/
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	inBounds := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int, int) bool

	dfs = func(r, c, curL int) bool {
		if curL >= len(word) {
			return false
		}

		if !inBounds(r, c) || visited[r][c] || board[r][c] != word[curL] {
			return false
		}

		if len(word)-1 == curL {
			return true
		}

		visited[r][c] = true
		for _, m := range moves {
			nR, nC := r+m[0], c+m[1]
			if dfs(nR, nC, curL+1) {
				return true
			}
		}
		visited[r][c] = false

		return false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] {
				if dfs(r, c, 0) {
					return true
				}
			}
		}
	}

	return false
}
