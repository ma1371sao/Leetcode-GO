func firstMissingPositive(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 1
    }
    slow := 0
    fast := 0
    for ; fast < n; fast++ {
        if nums[fast] >= 1 && nums[fast] <= n {
            nums[slow] = nums[fast]
            slow++
        }
    }
    for i := 0; i < slow; i++ {
        idx := abs(nums[i]) - 1
        if nums[idx] > 0 {
            nums[idx] = -nums[idx]
        }
    }
    for i := 0; i < slow; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    return slow + 1
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
