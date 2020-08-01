func maxLength(arr []string) int {
    uqChStrs := make([]string, 0)
    for _, s := range arr {
        if isUqChStr(s) {
            uqChStrs = append(uqChStrs, s)
        }
    }
    
    return finMaxLenUqChSubseq(uqChStrs, "", 0, 0)
}

func finMaxLenUqChSubseq(uqChStrs []string, curSubseq string, idx, maxLen int) int {
    if idx == len(uqChStrs) {        
        if isUqChStr(curSubseq) && len(curSubseq) > maxLen {
            return len(curSubseq)
        }
        return maxLen
    }
    
    maxLen = finMaxLenUqChSubseq(uqChStrs, curSubseq + uqChStrs[idx], idx + 1, maxLen)
    maxLen = finMaxLenUqChSubseq(uqChStrs, curSubseq, idx + 1, maxLen)
    return maxLen
}

func isUqChStr(str string) bool {
    chs := make(map[string]struct{})
    isUq := true
    for _, ch := range str {
        chStr := string(ch)
        if _, ok := chs[chStr]; !ok {
            chs[chStr] = struct{}{}
        } else {
            isUq = false
            break
        }
    }
    return isUq
}
