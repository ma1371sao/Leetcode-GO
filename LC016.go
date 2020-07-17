func Diff(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func threeSumClosest(nums []int, target int) int {
    if len(nums) < 3 {
        return 0
    }
    
    closest := nums[0] + nums[1] + nums[2]
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    for i := 0; i < len(nums) - 2; {
        for l, r := i + 1, len(nums) - 1; l < r; {
            sum := nums[i] + nums[l] + nums[r]
            if Diff(sum, target) < Diff(closest, target) {
                closest = sum
            }
            if sum < target {
                for l += 1; l < r && nums[l] == nums[l - 1]; l++ {}
            } else {
                for r -= 1; l < r && nums[r] == nums[r + 1]; r-- {}
            }
        }
        for i += 1; i < len(nums) - 2 && nums[i] == nums[i - 1]; i++ {}
    }
    return closest
}
