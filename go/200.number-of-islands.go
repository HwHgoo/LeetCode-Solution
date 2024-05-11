package leetcode

func numIslands(grid [][]byte) int {
	ans := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '1' {
				ans += 1
				walk(grid, x, y)
			}
		}
	}
	return ans
}

func walk(grid [][]byte, x, y int) {
	grid[y][x] = '0'
	// left
	if x > 0 && grid[y][x-1] == '1' {
		walk(grid, x-1, y)
	}

	// right
	if x < len(grid[0])-1 && grid[y][x+1] == '1' {
		walk(grid, x+1, y)
	}

	// down
	if y < len(grid)-1 && grid[y+1][x] == '1' {
		walk(grid, x, y+1)
	}

	// up
	if y > 0 && grid[y-1][x] == '1' {
		walk(grid, x, y-1)
	}
}
