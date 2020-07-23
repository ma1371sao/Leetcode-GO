func removeDuplicates(nums []int) int {
    nLen := 0
    for i := range nums {
        if i == 0 {
            nLen++
        } else {
            if nums[i] > nums[i - 1] {
                nums[nLen] = nums[i]
                nLen++
            }
        }
    }
    return nLen
}
