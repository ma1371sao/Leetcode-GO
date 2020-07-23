func searchRange(nums []int, target int) []int {
    res := make([]int, 2)
    res[0] = -1
    res[1] = -1
    if len(nums) == 0 {
        return res
    }
    
    l := 0
    r := len(nums) - 1
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    
    if nums[l] == target {
        res[0] = l
    } else {
        return res
    }
    
    l = 0
    r = len(nums) - 1
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] <= target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    
    res[1] = len(nums) - 1
    if nums[l] > target {
        res[1] = l - 1
    }
    return res
}
