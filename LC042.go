func trap(height []int) int {
    sizeWater := 0
    hIdxStack := make([]int, 0)
    for i := range height {
        if len(hIdxStack) > 0 && height[hIdxStack[0]] <= height[i] {
            for len(hIdxStack) > 0 {
                idx := len(hIdxStack) - 1
                sizeWater += height[hIdxStack[0]] - height[hIdxStack[idx]]
                hIdxStack = hIdxStack[:idx]
            }
        }
        hIdxStack = append(hIdxStack, i)
    }
    
    l := 0
    if len(hIdxStack) > 0 {
        l = hIdxStack[0]
    }
    hIdxStack = make([]int, 0)
    for i := len(height) - 1; i >= l; i-- {
        if len(hIdxStack) > 0 && height[hIdxStack[0]] <= height[i] {
            for len(hIdxStack) > 0 {
                idx := len(hIdxStack) - 1
                sizeWater += height[hIdxStack[0]] - height[hIdxStack[idx]]
                hIdxStack = hIdxStack[:idx]
            }
        }
        hIdxStack = append(hIdxStack, i)
    }
    return sizeWater
}
