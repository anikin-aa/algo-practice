package backtracking

func solveNQueens(n int) [][]string {
    cols, topLeftDiag, topRightDiag := make(map[int]bool), make(map[int]bool), make(map[int]bool)
    solutions := [][]string{}
    candBoard := initBoard(n, ".")

    var backtrack func(rowIdx int)
    backtrack = func(candRowIdx int) {
        if candRowIdx == n {
            solutions = append(solutions, append([]string{}, candBoard...))
            return
        }
        
        for i := 0; i < n ; i++ {
            colIdx,tldIdx, trdIdx := i, candRowIdx - i, candRowIdx + i
            if cols[colIdx] || topLeftDiag[tldIdx] || topRightDiag[trdIdx] { 
                continue
            }

            candBoard[candRowIdx] = replace(candBoard[candRowIdx], i, "Q")
            cols[colIdx], topLeftDiag[tldIdx], topRightDiag[trdIdx] = true,true,true

            backtrack(candRowIdx + 1)

            candBoard[candRowIdx] = replace(candBoard[candRowIdx], i, ".")
            cols[colIdx], topLeftDiag[tldIdx], topRightDiag[trdIdx] = false,false,false
        }
    }

    backtrack(0)
    return solutions
}

func initBoard(n int, val string) []string {
    board := make([]string,n)
    for i := range board {
       for j := 0; j < n ; j++ {
           board[i]+=val
       }
    }
    return board
}

func replace(str string, idx int, val string) string {
    toRuneSlice := []rune(str)
    toRuneSlice[idx] = []rune(val)[0]
    return string(toRuneSlice)
}
