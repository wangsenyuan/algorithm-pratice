package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}

	res := solve(n, m, E)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	fmt.Print(buf.String())
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

const INF = 1 << 60

func solve(n int, m int, E [][]int) []int64 {
	g := NewGraph(n, m*2)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u-1, v-1, int64(w))
		g.AddEdge(v-1, u-1, int64(w))
	}

	mn := make([]int64, n)
	dist := make([]int64, n)

	for i := 0; i < n; i++ {
		mn[i] = INF
		dist[i] = INF
	}

	//mn[0] = 0
	dist[0] = 0

	pq := make(PriorityQueue, 0, n)
	heap.Push(&pq, NewItem(0, 0))

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value
		if dist[u] < cur.priority {
			continue
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if mn[v] > g.val[i] {
				for j := g.nodes[v]; j > 0; j = g.next[j] {
					x := g.to[j]
					w := g.val[i] + g.val[j]
					if dist[x] > dist[u]+w*w {
						dist[x] = dist[u] + w*w
						heap.Push(&pq, NewItem(x, dist[x]))
					}
				}
				mn[v] = g.val[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if dist[i] == INF {
			dist[i] = -1
		}
	}
	return dist
}

func solve1(n int, m int, E [][]int) []int64 {
	g := NewGraph(n*51, 2*m*51)

	connect := func(u, v, w int) {
		g.AddEdge(u*51, v*51+w, 0)
		for was := 1; was <= 50; was++ {
			g.AddEdge(u*51+was, v*51, int64(was+w)*int64(was+w))
		}
	}

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		connect(u-1, v-1, w)
		connect(v-1, u-1, w)
	}

	dist := make([]int64, n*51)
	for i := 0; i < len(dist); i++ {
		dist[i] = INF
	}
	pq := make(PriorityQueue, 0, n)
	dist[0] = 0
	heap.Push(&pq, NewItem(0, 0))

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)

		u := cur.value

		if dist[u] < cur.priority {
			continue
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] > cur.priority+g.val[i] {
				dist[v] = cur.priority + g.val[i]
				heap.Push(&pq, NewItem(v, dist[v]))
			}
		}
	}

	res := make([]int64, n)

	for i := 0; i < n; i++ {
		if dist[51*i] == INF {
			res[i] = -1
		} else {
			res[i] = dist[i*51]
		}
	}

	return res
}

type Graph struct {
	count []int
	nodes []int
	next  []int
	to    []int
	val   []int64
	cur   int
}

func NewGraph(n int, e int) *Graph {
	count := make([]int, n)
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int64, e+10)
	return &Graph{count, nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v int, w int64) {
	g.count[u]++
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int   // The value of the item; arbitrary.
	priority int64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(v int, p int64) *Item {
	item := new(Item)
	item.value = v
	item.priority = p
	return item
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
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
func (pq *PriorityQueue) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
