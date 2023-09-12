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
	res := solve(n, E)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	fmt.Println(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const inf = 1 << 60

func solve(n int, E [][]int) []int64 {
	g := NewGraph(n, len(E))
	r := NewGraph(n, len(E))

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		r.AddEdge(v, u, w)
	}

	vs := make([]*Vertex, n)

	for i := 0; i < n; i++ {
		vs[i] = new(Vertex)
		vs[i].id = i
		vs[i].dist = inf
		vs[i].index = i
	}
	vs[0].dist = 0
	pq := make(PriorityQueue, n)

	copy(pq, vs)

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Vertex)
		if cur.dist == inf {
			break
		}
		u := cur.id

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vs[v].dist > cur.dist+int64(g.val[i]) {
				pq.update(vs[v], cur.dist+int64(g.val[i]))
			}
		}
	}

	// now we get the min time needed to reach it from 0

	pq = make(PriorityQueue, n)
	copy(pq, vs)
	for i := 0; i < n; i++ {
		vs[i].index = i
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Vertex)
		if cur.dist == inf {
			break
		}
		u := cur.id
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if vs[v].dist > cur.dist+int64(r.val[i]) {
				pq.update(vs[v], cur.dist+int64(r.val[i]))
			}
		}
	}

	ans := make([]int64, n-1)

	for i := 1; i < n; i++ {
		if vs[i].dist == inf {
			ans[i-1] = -1
		} else {
			ans[i-1] = vs[i].dist
		}
	}

	return ans
}

type Vertex struct {
	id    int
	dist  int64
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Vertex

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Vertex, dist int64) {
	if item.index < 0 {
		return
	}
	item.dist = dist
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
