func lengthOfLongestSubstring(s string) int {
    if len(s) < 2 {
        return len(s)
    }
    
    chIdx := make(map[byte]int)
    chIdx[s[0]] = 0
    start := 0
    maxLen := 1
    for i := 1; i < len(s); i++ {
        if idx, ok := chIdx[s[i]]; ok && idx >= start {
            start = idx + 1
        } else {
            curLen := i - start + 1
            if curLen > maxLen {
                maxLen = curLen
            }
        }
        chIdx[s[i]] = i
    }
    return maxLen
}
