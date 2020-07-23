func longestValidParentheses(s string) int {
    vldPthLen := make([]int, len(s) + 1)
    maxLen := 0
    for i := 1; i < len(s); i++ {
        if s[i] == '(' {
            vldPthLen[i + 1] = 0
        } else {
            if s[i - 1] == '(' {
                vldPthLen[i + 1] = vldPthLen[i - 1] + 2
            } else if vldPthLen[i] > 0 && i - vldPthLen[i] - 1 >= 0 && s[i - vldPthLen[i] - 1] == '(' {
                vldPthLen[i + 1] = vldPthLen[i] + 2 + vldPthLen[i - vldPthLen[i] - 1]
            }
        }
        
        if vldPthLen[i + 1] > maxLen {
            maxLen = vldPthLen[i + 1]
        }
    }
    return maxLen
}
