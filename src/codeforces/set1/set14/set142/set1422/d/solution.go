package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, m := readTwoNums(reader)
	s := readNNums(reader, 4)
	special := make([][]int, m)
	for i := 0; i < m; i++ {
		special[i] = readNNums(reader, 2)
	}
	res := solve(m, s[:2], s[2:], special)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const inf = 1 << 30

func solve(m int, s []int, f []int, special [][]int) int {
	// at most n + n
	// 那些通过 相同row/col组成的，可以组成一个group
	// 然后计算这些group的最短距离
	recs := make([]Record, m+1)
	for i := 0; i < m; i++ {
		recs[i] = Record{special[i][0], special[i][1], i}
	}
	recs[m] = Record{s[0], s[1], m}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].x < recs[j].x
	})

	g := NewGraph(m+1, 4*(m+1))
	for i := 0; i < m; i++ {
		a := recs[i].id
		b := recs[i+1].id
		dx := recs[i+1].x - recs[i].x
		g.AddEdge(a, b, dx)
		g.AddEdge(b, a, dx)
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].y < recs[j].y
	})

	for i := 0; i < m; i++ {
		a := recs[i].id
		b := recs[i+1].id
		dy := recs[i+1].y - recs[i].y
		g.AddEdge(a, b, dy)
		g.AddEdge(b, a, dy)
	}

	dist := make([]*Item, m+1)
	pq := make(PriorityQueue, m+1)
	for i := 0; i <= m; i++ {
		dist[i] = new(Item)
		dist[i].value = i
		if i < m {
			dist[i].priority = min(abs(special[i][0]-s[0]), abs(special[i][1]-s[1]))
		} else {
			dist[i].priority = 0
		}
		dist[i].index = i
		pq[i] = dist[i]
	}

	heap.Init(&pq)
	res := abs(s[0]-f[0]) + abs(s[1]-f[1])
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value
		if u < m {
			res = min(res, cur.priority+abs(special[u][0]-f[0])+abs(special[u][1]-f[1]))
		}
		var x, y int
		if u == m {
			x, y = s[0], s[1]
		} else {
			x, y = special[u][0], special[u][1]
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v < m && dist[v].priority > cur.priority+min(abs(x-special[v][0]), abs(y-special[v][1])) {
				pq.update(dist[v], cur.priority+min(abs(x-special[v][0]), abs(y-special[v][1])))
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type Record struct {
	x  int
	y  int
	id int
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
	if item.index < 0 {
		return
	}
	item.priority = priority
	heap.Fix(pq, item.index)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
