func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    curComb := make([]int, 0)
    return findCombination(candidates, target, res, 0, curComb)
}

func findCombination(candidates []int, target int, res [][]int, idx int, curComb []int) [][]int {
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
    
    res = findCombination(candidates, target, res, idx + 1, curComb)
    n := 1
    for candidates[idx] * n <= target {
        curComb = append(curComb, candidates[idx])
        res = findCombination(candidates, target - candidates[idx] * n, res, idx + 1, curComb)
        n++
    }
    return res
}
