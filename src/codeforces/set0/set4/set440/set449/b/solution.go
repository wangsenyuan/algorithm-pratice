package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	roads := make([][]int, m)

	for i := 0; i < m; i++ {
		roads[i] = readNNums(reader, 3)
	}
	s := make([]int, k)
	y := make([]int, k)

	for i := 0; i < k; i++ {
		s[i], y[i] = readTwoNums(reader)
	}

	res := solve(n, roads, s, y)

	fmt.Println(res)
}

const oo = 1 << 60

func solve(n int, roads [][]int, s []int, y []int) int {
	g := NewGraph(n, 2*len(roads)+len(s))

	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	for i := 0; i < len(s); i++ {
		j := s[i] - 1
		g.AddEdge(0, j, y[i])
	}
	pq := make(PriorityQueue, n)
	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.value = i
		it.priority = oo
		it.index = i
		items[i] = it
		pq[i] = it
	}
	items[0].priority = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.value
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if items[v].priority > items[u].priority+w {
				pq.update(items[v], items[u].priority+w)
			}
		}
	}

	deg := make([]int, n)
	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		u--
		v--
		if items[u].priority > items[v].priority {
			u, v = v, u
		}
		if items[u].priority+w == items[v].priority {
			deg[v]++
		}
	}

	for i := 0; i < len(s); i++ {
		j := s[i] - 1
		if items[j].priority == y[i] {
			deg[j]++
		}
	}

	var res int
	for i := 0; i < len(s); i++ {
		j := s[i] - 1
		if items[j].priority == y[i] {
			if deg[j] > 1 {
				deg[j]--
				res++
			}
		} else if items[j].priority < y[i] {
			res++
		}
	}

	return res
}

func solve1(n int, roads [][]int, s []int, y []int) int {
	g := NewGraph(n, 2*len(roads))

	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	trains := make([]Pair, len(s))
	for i := 0; i < len(s); i++ {
		trains[i] = Pair{y[i], s[i] - 1}
	}

	sort.Slice(trains, func(i, j int) bool {
		return trains[i].first < trains[j].first
	})

	pq := make(PriorityQueue, n)
	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.value = i
		it.priority = oo
		it.index = i
		items[i] = it
		pq[i] = it
	}
	items[0].priority = 0
	heap.Init(&pq)

	var res int
	for _, cur := range trains {
		id := cur.second
		// dist[id] > cur.first
		for pq.Len() > 0 && pq[0].priority <= cur.first {
			it := heap.Pop(&pq).(*Item)
			u := it.value
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if items[v].priority > items[u].priority+w {
					pq.update(items[v], items[u].priority+w)
				}
			}
		}

		if items[id].priority <= cur.first {
			res++
		} else if items[id].index >= 0 {
			pq.update(items[id], cur.first)
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
