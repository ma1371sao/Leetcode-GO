func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    if (m + n) % 2 == 1 {
        return float64(findKthNum((m + n) / 2 + 1, nums1, nums2))
    } else {
        return float64((findKthNum((m + n) / 2, nums1, nums2) + findKthNum((m + n) / 2 + 1, nums1, nums2))) / 2
    }
}

func findKthNum(k int, nums1 []int, nums2 []int) int {
    if len(nums1) < len(nums2) {
        return findKthNum(k, nums2, nums1)
    }
    
    if len(nums2) == 0 {
        return nums1[k - 1]
    }
    
    if k == 1 {
        if nums1[0] < nums2[0] {
            return nums1[0]
        } else {
            return nums2[0]
        }
    }
    
    var k1, k2 int
    if k / 2 > len(nums2) {
        k2 = len(nums2)
    } else {
        k2 = k / 2
    }
    k1 = k - k2
    
    if nums1[k1 - 1] == nums2[k2 - 1] {
        return nums1[k1 - 1]
    } else if nums1[k1 - 1] > nums2[k2 - 1] {
        return findKthNum(k1, nums1, nums2[k2:])
    } else {
        return findKthNum(k2, nums1[k1:], nums2)
    }
}
