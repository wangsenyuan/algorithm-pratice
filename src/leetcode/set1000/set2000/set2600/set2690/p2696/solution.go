package p2696

import "container/heap"

const INF = 2 * 1e9

func modifiedGraphEdges(n int, edges [][]int, source int, dest int, target int) [][]int {
	// 先建用正值边相连的节点，变成一个个comp

	g := NewGraph(n, len(edges)*2)

	type Pair struct {
		first  int
		second int
	}
	book := make(map[Pair]int)

	for i, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w < 0 {
			if u > v {
				u, v = v, u
			}
			book[Pair{u, v}] = i
			w = 1
		}

		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	pq := make(PriorityQueue, n)
	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.index = i
		it.value = INF
		items[i] = it
		pq[i] = it
	}

	items[source].value = 0
	heap.Init(&pq)

	dist0 := make([]int, n)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u, d := cur.id, cur.value
		dist0[u] = d
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if items[v].value > d+w {
				pq.update(items[v], d+w)
			}
		}
	}

	if dist0[dest] > target {
		return nil
	}

	expect := target - dist0[dest]

	dist1 := make([]int, n)

	for i := 0; i < n; i++ {
		items[i].value = INF
		heap.Push(&pq, items[i])
	}

	pq.update(items[source], 0)
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u, d := cur.id, cur.value
		dist1[u] = d
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			a, b := u, v
			if a > b {
				a, b = b, a
			}
			if j, ok := book[Pair{a, b}]; ok {
				z := dist0[v] + expect - dist1[u]
				if w < z {
					w = z
				}
				edges[j][2] = w
				delete(book, Pair{a, b})
			}
			nd := d + w
			if items[v].value > nd {
				pq.update(items[v], nd)
			}
		}
	}

	if dist1[dest] != target {
		return nil
	}

	for _, j := range book {
		edges[j][2] = INF
	}

	return edges
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	value := make([]int, e+1)
	return &Graph{nodes, next, to, value, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

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

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, dist int) {
	item.value = dist
	heap.Fix(pq, item.index)
}
