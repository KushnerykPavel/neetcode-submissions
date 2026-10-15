func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)

	var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)

	dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
		coord := [2]int{r, c}
		if visit[coord] || 
			r < 0 || c < 0 || 
			r == ROWS || c == COLS ||
			heights[r][c] < prevHeight {
				return 
			}

		visit[coord] = true

		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	for c := 0; c < COLS; c++ {
		dfs(0, c, pac, heights[0][c])
		dfs(ROWS-1, c, atl, heights[ROWS-1][c])
	}

	for r := 0; r < ROWS; r++ {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, COLS-1, atl, heights[r][COLS-1])
	}

	result := make([][]int, 0)
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			coord := [2]int{r,c}
			if pac[coord] && atl[coord] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result 
}
