package p2975

import (
	"container/heap"
)

const inf = 1 << 60

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	ids := make(map[string]int)
	root := NewNode()

	addWords := func(words []string) {
		for _, x := range words {
			if _, ok := ids[x]; ok {
				continue
			}
			ids[x] = len(ids)
			node := root
			for i := len(x) - 1; i >= 0; i-- {
				c := int(x[i] - 'a')
				if node.children[c] == nil {
					node.children[c] = NewNode()
				}
				node = node.children[c]
			}
			node.id = len(ids) - 1
		}
	}
	addWords(original)
	addWords(changed)

	m := len(ids)

	g := NewGraph(m, len(original))

	for i, x := range original {
		u := ids[x]
		v := ids[changed[i]]
		g.AddEdge(u, v, cost[i])
	}

	best := make([][]int, m)
	for i := 0; i < m; i++ {
		best[i] = make([]int, m)
		for j := 0; j < m; j++ {
			best[i][j] = inf
		}
	}

	items := make([]*Item, m)
	for i := 0; i < m; i++ {
		items[i] = new(Item)
		items[i].id = i
		items[i].value = inf
		items[i].index = i
	}

	bfs := func(start int) {
		pq := make(PriorityQueue, m)
		for i := 0; i < m; i++ {
			items[i].value = inf
			items[i].index = i
			pq[i] = items[i]
		}

		heap.Init(&pq)
		pq.update(items[start], 0)

		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(*Item)
			u := cur.id
			best[start][u] = cur.value
			if cur.value >= inf {
				break
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if items[v].value > cur.value+g.val[i] {
					pq.update(items[v], cur.value+g.val[i])
				}
			}
		}
	}

	for _, x := range original {
		bfs(ids[x])
	}

	n := len(source)
	// dp[i] 表示将source[0..i]与 target[0...i] 相等时的最小值
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		if source[i-1] == target[i-1] {
			dp[i] = dp[i-1]
		}
		a := root
		b := root
		for j := i - 1; j >= 0 && a != nil && b != nil; j-- {
			a = a.children[int(source[j]-'a')]
			b = b.children[int(target[j]-'a')]
			if a != nil && b != nil && a.id >= 0 && b.id >= 0 {
				dp[i] = min(dp[i], dp[j]+best[a.id][b.id])
			}
		}
	}
	if dp[n] >= inf {
		return -1
	}
	return int64(dp[n])
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
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, val int) {
	item.value = val
	heap.Fix(pq, item.index)
}

type Node struct {
	id       int
	children [26]*Node
}

func NewNode() *Node {
	node := new(Node)
	node.id = -1
	return node
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
