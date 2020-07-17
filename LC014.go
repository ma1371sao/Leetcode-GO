func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    for idx := range strs[0] {
        for i := 1; i < len(strs); i++ {
            if len(strs[i]) <= idx || strs[i][idx] != strs[0][idx] {
                return strs[0][:idx]
            }
        }
    }
    return strs[0]
}
