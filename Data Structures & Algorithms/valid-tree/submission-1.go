func validTree(n int, edges [][]int) bool {
    adj := make(map[int][]int, n)
	visited := make([]bool, n)

	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	var dfs func(v, parent int) bool 
	dfs = func(v, parent int) bool {
		visited[v] = true

		for _, e := range adj[v] {
			if e == parent {
				continue 
			}

			if visited[e] {
				return false 
			} 
			if !dfs(e, v) {
				return false
			}
		}

		return true
	}
	if !dfs(0, -1) {
		return false 
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
