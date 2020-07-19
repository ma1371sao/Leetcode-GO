func isValid(s string) bool {
    stack := make([]string, 0)
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, string(ch))
        } else {
            sLen := len(stack)
            if sLen == 0 {
                return false
            }

            if ch == ')' {
                if stack[sLen - 1] == "(" {
                    stack = stack[: sLen-1]
                } else {
                    return false
                }
            } else if ch == '}' {
                if stack[sLen - 1] == "{" {
                    stack = stack[: sLen-1]
                } else {
                    return false
                }
            } else if ch == ']' {
                if stack[sLen - 1] == "[" {
                    stack = stack[: sLen-1]
                } else {
                    return false
                }
            }
        }
    }
    
    if len(stack) > 0 {
        return false
    }
    return true
}
