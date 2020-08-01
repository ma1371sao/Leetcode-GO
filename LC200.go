func numIslands(grid [][]byte) int {
    num := 0
    visited := make([][]bool, len(grid))
    for r := range grid {
        visited[r] = make([]bool, len(grid[r]))
    }
    
    for r := range grid {
        for c := range grid[r] {
            if !visited[r][c] && grid[r][c] == '1' {
                dfs(grid, r, c, visited)
                num++
            }
        }
    }
    return num
}

func dfs(grid [][]byte, r, c int, visited [][]bool) {
    visited[r][c] = true
    numR := len(grid)
    numC := 0
    if numR > 0 {
        numC = len(grid[0])
    }
    
    if r > 0 && grid[r - 1][c] == '1' && !visited[r - 1][c] {
        dfs(grid, r - 1, c, visited)
    }
    
    if r < numR - 1 && grid[r + 1][c] == '1' && !visited[r + 1][c] {
        dfs(grid, r + 1, c, visited)
    }
    
    if c > 0 && grid[r][c - 1] == '1' && !visited[r][c - 1] {
        dfs(grid, r, c - 1, visited)
    }
    
    if c < numC - 1 && grid[r][c + 1] == '1' && !visited[r][c + 1] {
        dfs(grid, r, c + 1, visited)
    }
}
