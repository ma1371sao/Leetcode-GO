func nextPermutation(nums []int)  {
    if len(nums) == 0 {
        return
    }
    
    start := len(nums) - 1
    for ; start > 0 && nums[start - 1] >= nums[start]; start-- {}
    if start == 0 {
        sort.Slice(nums, func(i, j int) bool {
            return nums[i] < nums[j]
        })
        return
    }
    
    target := start
    for ; target < len(nums) && nums[target] > nums[start - 1]; target++ {}
    nums[start - 1], nums[target - 1] = nums[target - 1], nums[start - 1]
    sort.Slice(nums[start:], func(i, j int) bool {
        return nums[start:][i] < nums[start:][j]
    })
}
