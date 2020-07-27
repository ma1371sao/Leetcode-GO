func isValidSudoku(board [][]byte) bool {
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
            if _, ok := rowDigit[r][d]; ok {
                return false
            }
            rowDigit[r][d] = struct{}{}
            
            if _, ok := colDigit[c][d]; ok {
                return false
            }
            colDigit[c][d] = struct{}{}
            
            newR := r / 3
            newC := c / 3
            idx := newR * 3 + newC
            if _, ok := subboxDigit[idx][d]; ok {
                return false
            }
            subboxDigit[idx][d] = struct{}{}
        }
    }
    return true
}
