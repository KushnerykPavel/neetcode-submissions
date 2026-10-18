func findOrder(numCourses int, prerequisites [][]int) []int {
    adj := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		adj[prereq[0]] = append(adj[prereq[0]], prereq[1])
	}
	path := make([]int, 0, numCourses)
	state := make([]uint8, numCourses)

	var dfs func(v int) bool

	dfs = func(v int) bool {
		if state[v] == 1 {
			return false
		}
		if state[v] == 2 {
			return true
		}

		state[v] = 1
		for _, n := range adj[v] {
			if !dfs(n) {
				return false
			}
		}
		state[v] = 2
		path = append(path, v)

		return true
	}

	for i := 0; i< numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return path
}
