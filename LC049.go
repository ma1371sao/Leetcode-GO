func groupAnagrams(strs []string) [][]string {
    res := make([][]string, 0)
    wordIdx := make(map[string]int)
    for _, str := range strs {
        strBytes := []byte(str)
        sort.Slice(strBytes, func(i, j int) bool {
            return strBytes[i] < strBytes[j]
        })
        
        word := string(strBytes)
        if _, ok := wordIdx[word]; !ok {
            l := len(res)
            wordIdx[word] = l
            words := make([]string, 0)
            res = append(res, words)
        }
        
        res[wordIdx[word]] = append(res[wordIdx[word]], str)
    }
    return res
}
