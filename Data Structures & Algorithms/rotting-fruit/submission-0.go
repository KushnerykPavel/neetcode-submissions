func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	queue := make([][2]int, 0)
	fresh, time := 0, 0 

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
			if grid[r][c] == 1 {
				fresh++ 
			}
		}
	}

	for len(queue) > 0 && fresh > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, dir := range directions {
				newRow := row + dir[0]
				newCol := col + dir[1]

				if newRow >= 0 && newRow < rows &&
				newCol >= 0 && newCol < cols &&
				grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					queue = append(queue, [2]int{newRow, newCol})
					fresh--
				}
			}
		}
		time++
	}

	if fresh == 0 {
		return time
	}
	return -1
}
