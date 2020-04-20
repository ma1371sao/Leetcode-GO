func twoSum(nums []int, target int) []int {
    val2idx := make(map[int]int)
    for i, val := range nums {
        if idx, ok := val2idx[target - val]; ok {
            res := make([]int, 2)
            res[0] = idx
            res[1] = i
            return res
        }
        val2idx[val] = i
    }
    return nil
}
