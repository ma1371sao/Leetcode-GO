func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    visit := make(map[int]struct{})
    curPerm := make([]int, 0)
    return dfs(nums, curPerm, res, visit)
}

func dfs(nums, curPerm []int, res [][]int, visit map[int]struct{}) [][]int {
    if len(nums) == len(curPerm) {
        perm := make([]int, 0)
        for _, num := range curPerm {
            perm = append(perm, num)
        }
        res = append(res, perm)
        return res
    }
    
    visitVal := make(map[int]struct{})
    curLen := len(curPerm)
    for i, num := range nums {
        if _, ok := visit[i]; !ok {
            if _, ok := visitVal[num]; !ok {
                visit[i] = struct{}{}
                visitVal[num] = struct{}{}
                curPerm = append(curPerm, num)
                res = dfs(nums, curPerm, res, visit)
                curPerm = curPerm[:curLen]
                delete(visit, i)
            }
        }
    }
    return res
}
