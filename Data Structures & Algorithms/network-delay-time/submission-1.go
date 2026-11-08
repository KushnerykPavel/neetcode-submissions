type edge struct {
	to int
	weight int
}

type State struct {
    node int
    dist int
}

type MinHeap []State

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(State))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make([][]edge, n + 1)

	dist := make([]int, n+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[k] = 0
	for _, t := range times {
		from, to, weight := t[0], t[1], t[2]

		graph[from] = append(graph[from], edge{
			to: to,
			weight: weight,
		})
	}

	pq := &MinHeap{{node: k, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		if current.dist > dist[current.node] {
			continue
		}

		for _, edge := range graph[current.node] {
			newDist := current.dist + edge.weight

			if newDist < dist[edge.to] {
				dist[edge.to] = newDist

				heap.Push(pq, State{
					node: edge.to, 
					dist: newDist,
				})
			}
		}
	}

	result := 0

	for node := 1; node <= n; node++ {
		if dist[node] == math.MaxInt {
			return -1
		}

		if dist[node] > result {
			result = dist[node]
		}
	}

	return result
}
