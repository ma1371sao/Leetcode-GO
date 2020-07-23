func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    l := 0
    r := len(nums) - 1
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] == target {
            return mid
        }
        
        if nums[mid] < nums[r] {
            if target > nums[mid] && target <= nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }

        } else {
            if target >= nums[l] && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    
    if nums[l] == target {
        return l
    }
    return -1
}
