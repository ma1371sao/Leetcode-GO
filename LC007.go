func reverse(x int) int {
    positive := true
    if x < 0 {
        positive = false
        x = -x
    }
    
    s := strconv.Itoa(x)
    rs := reverseStr(s)
    if !positive {
        rs = "-" + rs
    }
    i32, err := strconv.ParseInt(rs, 10, 32)
    if err != nil {
        return 0
    }
    return int(i32)
}

func reverseStr(s string) string {
    rs := ""
    for _, ch := range s {
        rs = string(ch) + rs
    }
    return rs
}
