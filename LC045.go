func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    dp := make(map[int]int)
    numJump := 1
    dp[1] = nums[0]
    for dp[numJump] < len(nums) - 1 {
        numJump++
        maxIdx := dp[numJump - 1]
        for i := maxIdx; i >= 0; i-- {
            if nums[i] + i > maxIdx {
                maxIdx = nums[i] + i
            }
        }
        dp[numJump] = maxIdx
    }
    return numJump
}
