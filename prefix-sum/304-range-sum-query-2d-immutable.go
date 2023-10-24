package main

/*
	O(nm) - space, O(mn) - time
*/
type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	sum := make([][]int, rows+1)
	for i := range sum {
		sum[i] = make([]int, cols+1)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			sum[r+1][c+1] = sum[r+1][c] + sum[r][c+1] - sum[r][c] + matrix[r][c]
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
	return this.sum[r1][c1] + this.sum[r2+1][c2+1] - this.sum[r1][c2+1] - this.sum[r2+1][c1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
