func maxArea(height []int) int {
    maxArea := 0
    l := 0
    r := len(height) - 1
    for l < r {
        area := min(height[l], height[r]) * (r - l)
        if maxArea < area {
            maxArea = area
        }
        if height[l] < height[r] {
            for l += 1; l < r && height[l] < height[l - 1]; l++ {}
        } else if height[r] < height[l] {
            for r -= 1; r > l && height[r] < height[r + 1]; r-- {}
        } else {
            for l += 1; l < r && height[l] < height[l - 1]; l++ {}
            for r -= 1; r > l && height[r] < height[r + 1]; r-- {}
        }
    }
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
