func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    if len(nums) < 3 {
        return res
    }
    
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    for i := 0; i < len(nums) - 2; {
        for l, r := i + 1, len(nums) - 1; l < r; {
            if nums[i] + nums[l] + nums[r] < 0 {
                for l += 1; l < r && nums[l] == nums[l - 1]; l++ {}
            } else if nums[i] + nums[l] + nums[r] > 0 {
                for r -= 1; l < r && nums[r] == nums[r + 1]; r-- {}
            } else {
                triplet := make([]int, 3)
                triplet[0] = nums[i]
                triplet[1] = nums[l]
                triplet[2] = nums[r]
                res = append(res, triplet)
                
                for l += 1; l < r && nums[l] == nums[l - 1]; l++ {}
                for r -= 1; l < r && nums[r] == nums[r + 1]; r-- {}
            }
        }
        for i += 1; i < len(nums) - 2 && nums[i] == nums[i - 1]; i++ {}
    }
    return res
}
