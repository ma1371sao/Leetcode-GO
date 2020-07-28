func combinationSum2(candidates []int, target int) [][]int {
    candNoDup := make([]int, 0)
    valCnt := make(map[int]int)
    for _, val := range candidates {
        if _, ok := valCnt[val]; ok {
            valCnt[val]++
        } else {
            valCnt[val] = 1
            candNoDup = append(candNoDup, val)
        }
    }
    
    res := make([][]int, 0)
    curComb := make([]int, 0)
    return findComb(candNoDup, valCnt, target, 0, curComb, res)
}

func findComb(candidates []int, valCnt map[int]int, target, idx int, curComb []int, res [][]int) [][]int {
    if target == 0 {
        comb := make([]int, 0)
        for _, val := range curComb {
            comb = append(comb, val)
        }
        res = append(res, comb)
        return res
    }
    
    if idx >= len(candidates) {
        return res
    }
    
    res = findComb(candidates, valCnt, target, idx + 1, curComb, res)
    for i := 0; i < valCnt[candidates[idx]]; i++ {
        if candidates[idx] <= target {
            target -= candidates[idx]
            curComb = append(curComb, candidates[idx])
            res = findComb(candidates, valCnt, target, idx + 1, curComb, res)
        }
    }
    return res
}
