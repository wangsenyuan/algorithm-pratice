package p3120

import "container/heap"

const oo = 1 << 60

func findAnswer(n int, edges [][]int) []bool {
	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.value = oo
		it.index = i
		items[i] = it
	}

	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	bfs := func(x int) []int {
		pq := make(PriorityQueue, n)
		for i := 0; i < n; i++ {
			items[i].value = oo
			items[i].index = i
			pq[i] = items[i]
		}
		pq[x].value = 0
		heap.Init(&pq)

		dist := make([]int, n)
		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			dist[it.id] = it.value
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if items[v].index < 0 {
					continue
				}
				if items[v].value > items[u].value+g.val[i] {
					pq.update(items[v], items[u].value+g.val[i])
				}
			}
		}

		return dist
	}

	dist1 := bfs(0)
	dist2 := bfs(n - 1)

	expect := dist1[n-1]

	ans := make([]bool, len(edges))

	for i, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if dist1[u] < oo && dist2[v] < oo && dist1[u]+dist2[v]+w == expect {
			ans[i] = true
		} else if dist1[v] < oo && dist2[u] < oo && dist1[v]+dist2[u]+w == expect {
			ans[i] = true
		}
	}

	return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	id    int
	value int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.index)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
