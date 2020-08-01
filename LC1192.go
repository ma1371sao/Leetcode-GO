func criticalConnections(n int, connections [][]int) [][]int {
    criticalConn := make([][]int, 0)
    adjNodes := make(map[int]map[int]struct{})
    for i := 0; i < n; i++ {
        adjNodes[i] = make(map[int]struct{})
    }
    for _, conn := range connections {
        adjNodes[conn[0]][conn[1]] = struct{}{}
        adjNodes[conn[1]][conn[0]] = struct{}{}
    }
    
    nodeDeep := make(map[int]int)
    visited := make(map[int]struct{})
    _, criticalConn = dfs(adjNodes, 0, 0, -1, criticalConn, nodeDeep, visited)
    return criticalConn
}

func dfs(adjNodes map[int]map[int]struct{}, node, deep, parent int, criticalConn [][]int, nodeDeep map[int]int, 
         visited map[int]struct{}) (int, [][]int) {
    
    nodeDeep[node] = deep
    visited[node] = struct{}{}
    lowestDeep := deep
    for adjNode := range adjNodes[node] {
        if _, ok := visited[adjNode]; !ok {
            var childLowestDeep int
            childLowestDeep, criticalConn = dfs(adjNodes, adjNode, deep + 1, node, criticalConn, nodeDeep, visited)
            if childLowestDeep > deep {
                conn := []int{node, adjNode}
                criticalConn = append(criticalConn, conn)
            }
            if childLowestDeep < lowestDeep {
                lowestDeep = childLowestDeep
            }
            
        } else if adjNode != parent && nodeDeep[adjNode] < lowestDeep {
            lowestDeep = nodeDeep[adjNode]
        }
    }
    return lowestDeep, criticalConn
}
