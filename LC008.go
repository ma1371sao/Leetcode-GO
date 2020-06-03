func myAtoi(str string) int {
    if str == "" {
        return 0
    }
    
    s := ""
    for idx, ch := range str {
        if ch != ' ' {
            s = str[idx:]
            break
        }
    }
    
    if s == "" {
        return 0
    }
    idx := 0
    if s[0] == '+' || s[0] == '-' {
        idx = 1
    }
    for ; idx < len(s); idx++ {
        if s[idx] < '0' || s[idx] > '9' {
            break
        }
    }
    s = s[:idx]
    if s == "" || s == "+" || s == "-" {
        return 0
    }
    i32, err := strconv.ParseInt(s, 10, 32)
    if err != nil {
        if s[0] != '-' {
            return (1<<31) - 1
        }
        return (1<<31) * -1
    }
    return int(i32)
}
