func canJump(nums []int) bool {
    minJump := len(nums) - 1
    for i := minJump - 1; i >= 0; i-- {
        if i + nums[i] >= minJump {
            minJump = i
        }
    }
    
    if minJump == 0 {
        return true
    }
    return false
}
