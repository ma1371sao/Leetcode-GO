func reorderLogFiles(logs []string) []string {
    slow, fast := len(logs) - 1, len(logs) - 1
    for fast >= 0 {
        if isDigLog(logs[fast]) {
            logs[fast], logs[slow] = logs[slow], logs[fast]
            slow--
        }
        fast--
    }
    
    if slow < 0 {
        return logs
    }
    
    sort.Slice(logs[:slow + 1], func(i, j int) bool {
        log1 := logs[:slow + 1][i]
        log2 := logs[:slow + 1][j]
        words1 := strings.Split(log1, " ")
        words2 := strings.Split(log2, " ")
        part1 := strings.Join(words1[1:], " ")
        part2 := strings.Join(words2[1:], " ")
        return part1 <= part2
    })
    return logs
}

func isDigLog(log string) bool {
    words := strings.Split(log, " ")
    for _, ch := range words[1] {
        if ch < '0' || ch > '9' {
            return false
        }
    }
    return true
}
