package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n, m, k := readThreeNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, k, edges)
}

func solve(n int, k int, edges [][]int) int {

	get := func(u int, v int) int {
		u, v = min(u, v), max(u, v)
		return u*n + v
	}

	dist := make(map[int]*Item)

	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = get(i, i)
		it.priority = 0
		dist[it.id] = it
	}

	slices.SortFunc(edges, func(a, b []int) int {
		return a[2] - b[2]
	})

	if len(edges) > k {
		edges = edges[:k]
	}

	pq := make(PriorityQueue, 0, len(edges))

	g := NewGraph(n, 2*len(edges))

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		id := get(u-1, v-1)
		it := new(Item)
		it.id = id
		it.priority = w
		dist[id] = it
		heap.Push(&pq, it)
		g.AddEdge(u-1, v-1, w)
		g.AddEdge(v-1, u-1, w)
	}

	update := func(u int, p int, d int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			id := get(p, v)
			if it, ok := dist[id]; ok && it.priority > d+w {
				pq.update(it, d+w)
			} else if !ok {
				it = new(Item)
				it.id = id
				it.priority = d + w
				dist[id] = it
				heap.Push(&pq, it)
			}
		}
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		x, y := it.id/n, it.id%n

		k--
		if k == 0 {
			return it.priority
		}

		update(x, y, it.priority)
		update(y, x, it.priority)
	}

	return -1
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
	id       int // The value of the item; arbitrary.
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
