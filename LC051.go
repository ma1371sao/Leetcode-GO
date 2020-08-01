func solveNQueens(n int) [][]string {
    res := make([][]string, 0)
    qRowCol := make(map[int]int)
    qCol := make(map[int]struct{})
    yMinusX := make(map[int]struct{})
    yPlusX := make(map[int]struct{})
    return dfs(res, qRowCol, qCol, yMinusX, yPlusX, 0, n)
}

func dfs(res [][]string, qRowCol map[int]int, qCol, yMinusX, yPlusX map[int]struct{}, row, n int) [][]string {
    if row == n {
        board := make([]string, 0)
        for r := 0; r < n; r++ {
            boardRow := ""
            for c := 0; c < n; c++ {
                if qRowCol[r] == c {
                    boardRow += "Q"
                } else {
                    boardRow += "."
                }
            }
            board = append(board, boardRow)
        }
        res = append(res, board)
        return res
    }
    
    for c := 0; c < n; c++ {
        if _, ok := qCol[c]; !ok {
            ymx := row - c
            if _, ok := yMinusX[ymx]; !ok {
                ypx := row + c
                if _, ok := yPlusX[ypx]; !ok {
                    yMinusX[ymx] = struct{}{}
                    yPlusX[ypx] = struct{}{}
                    qRowCol[row] = c
                    qCol[c] = struct{}{}
                    res = dfs(res, qRowCol, qCol, yMinusX, yPlusX, row + 1, n)
                    delete(qCol, c)
                    delete(yMinusX, ymx)
                    delete(yPlusX, ypx)
                }
            }
        }
    }
    return res
}
