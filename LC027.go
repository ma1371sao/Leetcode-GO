func removeElement(nums []int, val int) int {
    nLen := 0
    for i := range nums {
        if nums[i] != val {
            nums[nLen] = nums[i]
            nLen++
        }
    }
    return nLen
}
