func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)

	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	state := make([]uint8, numCourses)

	var dfs func(int) bool 
	dfs = func(v int) bool {
		// cycle
		if state[v] == 1 {
			return false
		}
		// fully checked
		if state[v] == 2 {
			return true
		}

		state[v] = 1
		for _, p := range graph[v] {
			if !dfs(p) {
				return false
			}
		}
		state[v] = 2
		return true
	}

	for v := 0; v < numCourses; v++ {
		if !dfs(v) {
			return false
		}
	}
	return true
}
