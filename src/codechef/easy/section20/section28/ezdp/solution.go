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

	n, m, q := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	F := make([][]int, n)
	for i := 0; i < n; i++ {
		F[i] = readNNums(reader, 3)
	}
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 4)
	}
	res := solve(n, E, F, Q)
	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(N int, E [][]int, F [][]int, Q [][]int) []int64 {
	g := NewGraph(N, len(E)*2)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	dist1 := dijkrastra(N, g, 0)
	distn := dijkrastra(N, g, N-1)

	fs := make([]*Node, len(F))
	set := make(PQ, len(F))

	for i := 0; i < len(F); i++ {
		x := int64(F[i][0]) + int64(F[i][1]) + int64(F[i][2])
		x *= int64(N - 1)
		x += dist1[i]
		x += distn[i]
		cur := NewNode(i, x)
		cur.index = i
		fs[i] = cur
		set[i] = cur
	}

	heap.Init(&set)

	res := make([]int64, len(Q))

	for j, cur := range Q {
		i, a, b, c := cur[0], cur[1], cur[2], cur[3]
		i--
		x := (int64(a) + int64(b) + int64(c)) * int64(N-1)
		x += dist1[i]
		x += distn[i]

		fs[i].priority = x
		heap.Fix(&set, fs[i].index)
		res[j] = set[0].priority
	}
	return res
}

const INF = 1 << 60

func dijkrastra(n int, g *Graph, start int) []int64 {
	nodes := make([]*Node, n)
	pq := make(PQ, n)

	for i := 0; i < n; i++ {
		cur := NewNode(i, INF)
		cur.index = i
		nodes[i] = cur
		pq[i] = cur
	}
	nodes[start].priority = 0

	res := make([]int64, n)

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node)
		res[cur.value] = cur.priority
		u := cur.value
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.cost[i]
			if nodes[v].priority > cur.priority+int64(w) {
				nodes[v].priority = cur.priority + int64(w)
				heap.Fix(&pq, nodes[v].index)
			}
		}
	}
	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cost  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cost = make([]int, e+1)
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

type Node struct {
	value    int
	priority int64
	index    int
}

func NewNode(v int, p int64) *Node {
	n := new(Node)
	n.value = v
	n.priority = p
	n.index = -1
	return n
}

type PQ []*Node

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].priority < this[j].priority
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *PQ) Push(x interface{}) {
	i := x.(*Node)
	i.index = len(*this)
	*this = append(*this, i)
}

func (this *PQ) Pop() interface{} {
	old := *this
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*this = old[:n-1]
	return ret
}
