func permute(nums []int) [][]int {
    res := make([][]int, 0)
    visit := make(map[int]struct{})
    curPerm := make([]int, 0)
    return dfs(nums, res, visit, curPerm)
}

func dfs(nums []int, res [][]int, visit map[int]struct{}, curPerm []int) [][]int {
    if len(curPerm) == len(nums) {
        perm := make([]int, 0)
        for _, num := range curPerm {
            perm = append(perm, num)
        }
        res = append(res, perm)
        return res
    }
    
    curLen := len(curPerm)
    for _, num := range nums {
        if _, ok := visit[num]; !ok {
            curPerm = append(curPerm, num)
            visit[num] = struct{}{}
            res = dfs(nums, res, visit, curPerm)
            delete(visit, num)
            curPerm = curPerm[:curLen]
        }
    }
    return res
}
