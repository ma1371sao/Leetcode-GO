func canReach(arr []int, start int) bool {
    reachIdx := make(map[int]struct{})
    reachIdx[start] = struct{}{}
    newReach := make(map[int]struct{})
    newReach[start] = struct{}{}
    
    for len(newReach) > 0 {
        newnewReach := make(map[int]struct{})
        for idx := range newReach {
            if idx + arr[idx] < len(arr) {
                if _, ok := reachIdx[idx + arr[idx]]; !ok {
                    reachIdx[idx + arr[idx]] = struct{}{}
                    newnewReach[idx + arr[idx]] = struct{}{}
                }
            }
            
            if idx - arr[idx] >= 0 {
                if _, ok := reachIdx[idx - arr[idx]]; !ok {
                    reachIdx[idx - arr[idx]] = struct{}{}
                    newnewReach[idx - arr[idx]] = struct{}{}
                }
            }
        }
        
        newReach = newnewReach
    }
    
    for i := range arr {
        if _, ok := reachIdx[i]; ok && arr[i] == 0 {
            return true
        }
    }
    return false
}
