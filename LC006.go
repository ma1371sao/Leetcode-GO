func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    res := ""
    for r := 0; r < numRows; r++ {
        idx := r
        downward := false

        for idx < len(s) {
            res += string(s[idx])
            if r == 0 {
                downward = true
            } else if r == numRows - 1 {
                downward = false
            } else {
                downward = (!downward)
            }

            if downward {
                idx += 2 * (numRows - 1 - r)
            } else {
                idx += r * 2
            }
        }
    }
    return res
}
