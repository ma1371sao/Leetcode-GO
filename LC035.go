func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    
    l := 0
    r := len(nums) - 1
    if nums[r] < target {
        return r + 1
    }
    if nums[l] > target {
        return 0
    }

    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] == target {
            return mid
        }
        
        if nums[mid] > target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}
