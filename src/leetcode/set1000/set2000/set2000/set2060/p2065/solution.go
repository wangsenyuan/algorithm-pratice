package p2065

import "container/heap"

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	//maxTime <= 100, values[i][2] >= 10, so maxTime / values[i][2] <= 10, and at most a 10 nodes path,
	//包括0的一个10节点连通集合，是否能在maxTime内返回, 获得最大值
	n := len(values)

	g := NewGraph(n, len(edges)*2)

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	vis := make([]bool, n)

	var ans int

	var dfs func(u, t, value int)

	dfs = func(u, t, value int) {
		if u == 0 {
			ans = max(ans, value)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.cost[i]
			if t+w <= maxTime {
				if !vis[v] {
					vis[v] = true
					dfs(v, t+w, value+values[v])
					vis[v] = false
				} else {
					dfs(v, t+w, value)
				}
			}
		}
	}
	vis[0] = true
	dfs(0, 0, values[0])

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cost  []int
	cur   int
}

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cost = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = w
}

type Item struct {
	value    int
	priority int
	index    int
}

func (this Item) Less(that Item) bool {
	return this.priority > that.priority
}

type PriroityQueue []*Item

func (pq PriroityQueue) Len() int {
	return len(pq)
}

func (pq PriroityQueue) Less(i, j int) bool {
	return pq[i].Less(*pq[j])
}

func (pq PriroityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriroityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriroityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}

func (pq *PriroityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
