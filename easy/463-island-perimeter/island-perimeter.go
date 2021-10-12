func islandPerimeter(grid [][]int) int {
    var perimeter int

    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[y]); x++ {
            if grid[y][x] < 1 { continue }
            
            if x-1 < 0 || grid[y][x-1] < 1 {
                perimeter += 1
            }
            
            if x+1 >= len(grid[y]) || grid[y][x+1] < 1 {
                perimeter += 1
            }
            
            if y-1 < 0 || grid[y-1][x] < 1 {
                perimeter += 1
            }
            
            if y+1 >= len(grid) || grid[y+1][x] < 1 {
                perimeter += 1
            }
        }
    }

    return perimeter
}

