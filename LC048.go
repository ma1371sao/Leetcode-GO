func rotate(matrix [][]int)  {
    n := len(matrix)
    for i, j := 0, n - 1; i < j; {
        for c := 0; c < n; c++ {
            tmp := matrix[i][c]
            matrix[i][c] = matrix[j][c]
            matrix[j][c] = tmp
        }
        i++
        j--
    }
    
    for r := 0; r < n; r++ {
        for c := r + 1; c < n; c++ {
            tmp := matrix[r][c]
            matrix[r][c] = matrix[c][r]
            matrix[c][r] = tmp
        }
    }
}
