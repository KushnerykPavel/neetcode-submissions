func solve(board [][]byte) {
    ROWS, COLS := len(board), len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int ) {
		if r < 0 || c < 0 ||
		 r >= ROWS || c >= COLS || board[r][c] != 'O' {
				return
		}
		board[r][c] = 'T'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if board[r][c] == 'O' && (
				r == 0 || r == ROWS - 1 || c == 0 || c == COLS - 1) {
				dfs(r, c)
			}
		}
	}

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if board[r][c] != 'T' {
				board[r][c] = 'X'
			}
		}
	}

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if board[r][c] == 'T' {
				board[r][c] = 'O'
			}
		}
	}
}
