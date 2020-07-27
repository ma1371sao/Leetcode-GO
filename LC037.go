func solveSudoku(board [][]byte)  {
    rowDigit := make(map[int]map[int]struct{})
    colDigit := make(map[int]map[int]struct{})
    subboxDigit := make(map[int]map[int]struct{})
    for i := 0; i < 9; i++ {
        rowDigit[i] = make(map[int]struct{})
        colDigit[i] = make(map[int]struct{})
        subboxDigit[i] = make(map[int]struct{})
    }
    
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                continue
            }
            
            d, _ := strconv.Atoi(string(board[r][c]))
            rowDigit[r][d] = struct{}{}
            colDigit[c][d] = struct{}{}
            newR := r / 3
            newC := c / 3
            idx := newR * 3 + newC
            subboxDigit[idx][d] = struct{}{}
        }
    }
    
    fillDigit(board, rowDigit, colDigit, subboxDigit, 0, 0)
}

func fillDigit(board [][]byte, rowDigit, colDigit, subboxDigit map[int]map[int]struct{}, row, col int) bool {
    for r := row; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                for d := 1; d <= 9; d++ {
                    if _, ok := rowDigit[r][d]; !ok {
                        if _, ok := colDigit[c][d]; !ok {
                            newR := r / 3
                            newC := c / 3
                            idx := newR * 3 + newC
                            if _, ok := subboxDigit[idx][d]; !ok {
                                board[r][c] = fmt.Sprint(d)[0]
                                rowDigit[r][d] = struct{}{}
                                colDigit[c][d] = struct{}{}
                                subboxDigit[idx][d] = struct{}{}
                                if fillDigit(board, rowDigit, colDigit, subboxDigit, r, c) {
                                    return true
                                }
                                board[r][c] = '.'
                                delete(rowDigit[r], d)
                                delete(colDigit[c], d)
                                delete(subboxDigit[idx], d)
                            }
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}
