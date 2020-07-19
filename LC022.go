func generateParenthesis(n int) []string {
    res := make([]string, 0)
    if n == 0 {
        return res
    }
    return genParenthesisStr(n, n, 0, "", res)
}

func genParenthesisStr(numFront, numBack, numUnPair int, parenthesisStr string, res []string) []string {
    if numFront == 0 && numBack == 0 && numUnPair == 0 {
        res = append(res, parenthesisStr)
        return res
    }
    
    if numFront > 0 {
        res = genParenthesisStr(numFront - 1, numBack, numUnPair + 1, parenthesisStr + "(", res)
    }
    if numUnPair > 0 {
        res = genParenthesisStr(numFront, numBack - 1, numUnPair - 1, parenthesisStr + ")", res)
    }
    return res
}
