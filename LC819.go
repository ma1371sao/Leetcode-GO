func mostCommonWord(paragraph string, banned []string) string {
    maxFreq := 0
    maxFreqWord := ""
    banWords := make(map[string]struct{})
    for _, word := range banned {
        banWords[word] = struct{}{}
    }
    
    word := ""
    wordsFreq := make(map[string]int)
    for _, ch := range paragraph {
        if ch >= 'a' && ch <= 'z' {
            word += string(ch)
        } else if ch >= 'A' && ch <= 'Z' {
            word += string(ch - 'A' + 'a')
        } else {
            if len(word) > 0 {
                if _, ok := banWords[word]; !ok {
                    wordsFreq[word]++
                    if maxFreq < wordsFreq[word] {
                        maxFreq = wordsFreq[word]
                        maxFreqWord = word
                    }
                }
            }
            word = ""
        }
    }
    if len(word) > 0 {
        if _, ok := banWords[word]; !ok {
            wordsFreq[word]++
            if maxFreq < wordsFreq[word] {
                maxFreq = wordsFreq[word]
                maxFreqWord = word
            }
        }
    }

    return maxFreqWord
}
