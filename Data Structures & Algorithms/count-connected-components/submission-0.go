type DSU struct {
	parent []int
	rank []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = 1
	}
	return dsu
}

func (dsu *DSU) Find(node int) int {
	cur := node 
	for cur != dsu.parent[cur] {
		dsu.parent[cur] = dsu.parent[dsu.parent[cur]]
		cur = dsu.parent[cur]
	}
	return cur
}

func (dsu *DSU) Union(u, v int) bool {
	pu := dsu.Find(u)
	pv := dsu.Find(v)
	if pu == pv {
		return false
	}

	if dsu.rank[pv] > dsu.rank[pu] {
		pu, pv = pv, pu
	}
	dsu.parent[pv] = pu
	dsu.rank[pu] += dsu.rank[pv]
	return true
}

func countComponents(n int, edges [][]int) int {
    dsu := NewDSU(n)
	res := n 
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if dsu.Union(u, v) {
			res--
		}
	}
	return res
}
