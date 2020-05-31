import "fmt"

func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    
    r := len(s) - 1
    lpsLen := 1
    lps := string(s[0])
    for m := r / 2; m + 1 <= r; m++ {
        i := 0
        for ; (m + i <= r) && (m - i >= 0) && s[m + i] == s[m - i]; i++ {}
        i--
        if i * 2 + 1 > lpsLen {
            lpsLen = i * 2 + 1
            lps = s[m - i:m + i + 1]
            fmt.Println(lps)
        }
        
        if s[m] == s[m + 1] {
            i := 0
            for ; (m + 1 + i <= r) && (m - i >= 0) && s[m + 1 + i] == s[m - i]; i++ {}
            i--
            if i * 2 + 2 > lpsLen {
                lpsLen = i * 2 + 2
                lps = s[m - i:m + 1 + i + 1]
                fmt.Println(lps)
            }
        }
    }
    
    for m := r / 2 - 1; m >= 0; m-- {
        i := 0
        for ; (m + i <= r) && (m - i >= 0) && s[m + i] == s[m - i]; i++ {}
        i--
        if i * 2 + 1 > lpsLen {
            lpsLen = i * 2 + 1
            lps = s[m - i:m + i + 1]
            fmt.Println(lps)
        }
        
        if s[m] == s[m + 1] {
            i := 0
            for ; (m + 1 + i <= r) && (m - i >= 0) && s[m + 1 + i] == s[m - i]; i++ {}
            i--
            if i * 2 + 2 > lpsLen {
                lpsLen = i * 2 + 2
                lps = s[m - i : m + 1 + i + 1]
                fmt.Println(lps)
            }
        }
    }
    
    return lps
}
