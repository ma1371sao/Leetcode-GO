func totalNQueens(n int) int {
    qRowCol := make(map[int]int)
    qCol := make(map[int]struct{})
    yMinusX := make(map[int]struct{})
    yPlusX := make(map[int]struct{})
    return dfs(qRowCol, qCol, yMinusX, yPlusX, 0, n, 0)
}

func dfs(qRowCol map[int]int, qCol, yMinusX, yPlusX map[int]struct{}, row, n, count int) int {
    if row == n {
        return count + 1
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
                    count = dfs(qRowCol, qCol, yMinusX, yPlusX, row + 1, n, count)
                    delete(qCol, c)
                    delete(yMinusX, ymx)
                    delete(yPlusX, ypx)
                }
            }
        }
    }
    return count
}
