package main

/*
https://leetcode.com/problems/matrix-block-sum/description/
*/

func matrixBlockSum(mat [][]int, k int) [][]int {
	rows, cols := len(mat), len(mat[0])

	sum := make([][]int, rows+1)
	for i := range sum {
		sum[i] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			sum[r+1][c+1] = sum[r+1][c] + sum[r][c+1] - sum[r][c] + mat[r][c]
		}
	}

	res := make([][]int, rows)
	for i := range res {
		res[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			r1, c1 := max(0, r-k), max(0, c-k)
			r2, c2 := min(rows-1, r+k), min(cols-1, c+k)

			r1, c1 = r1+1, c1+1
			r2, c2 = r2+1, c2+1
			res[r][c] = sum[r2][c2] - sum[r2][c1-1] - sum[r1-1][c2] + sum[r1-1][c1-1]
		}
	}

	return res
}
