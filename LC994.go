type Orange struct {
    r, c, time int
}

func orangesRotting(grid [][]int) int {
    numR := len(grid)
    numC := 0
    if numR > 0 {
        numC = len(grid[0])
    }
    
    q := make([]*Orange, 0)
    timeRotten := 0
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 2 {
                orange := &Orange{
                    r:    r,
                    c:    c,
                    time: 0,
                }
                q = append(q, orange)
            }
        }
    }
    
    for len(q) > 0 {
        orange := q[0]
        q = q[1:]
        if orange.time > timeRotten {
            timeRotten = orange.time
        }
        
        if orange.r > 0 && grid[orange.r - 1][orange.c] == 1 {
            grid[orange.r - 1][orange.c] = 2
            up := &Orange{
                r:    orange.r - 1,
                c:    orange.c,
                time: orange.time + 1,
            }
            q = append(q, up)
        }
        
        if orange.r < numR - 1 && grid[orange.r + 1][orange.c] == 1 {
            grid[orange.r + 1][orange.c] = 2
            down := &Orange{
                r:    orange.r + 1,
                c:    orange.c,
                time: orange.time + 1,
            }
            q = append(q, down)
        }
        
        if orange.c > 0 && grid[orange.r][orange.c - 1] == 1 {
            grid[orange.r][orange.c - 1] = 2
            left := &Orange{
                r:    orange.r,
                c:    orange.c - 1,
                time: orange.time + 1,
            }
            q = append(q, left)
        }
        
        if orange.c < numC - 1 && grid[orange.r][orange.c + 1] == 1 {
            grid[orange.r][orange.c + 1] = 2
            right := &Orange{
                r:    orange.r,
                c:    orange.c + 1,
                time: orange.time + 1,
            }
            q = append(q, right)
        }
    }
    
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 1 {
                return -1
            }
        }
    }
    return timeRotten
}
