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

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	Q := make([][]int, k)
	for i := 0; i < k; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, E, a, Q)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, E [][]int, a []int, Q [][]int) []int {

	g := NewGraph(n, 2*len(E))
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = new(Node)
		nodes[i].id = i
	}
	fa := make([]int, n)
	sum := make([]int, n)
	sons := make([]PQ, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		nodes[u].value++
		sons[u] = make(PQ, 0, 1)
		sum[u] = a[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				fa[v] = u
				dfs(u, v)
				nodes[u].value += nodes[v].value
				sum[u] += sum[v]
				heap.Push(&sons[u], nodes[v])
			}
		}
	}

	dfs(0, 0)

	update := func(u int) {
		if len(sons[u]) == 0 {
			// a leaf node
			return
		}
		first := heap.Pop(&sons[u]).(*Node)
		v := first.id
		p := fa[u]
		newValue := nodes[u].value - nodes[v].value
		sons[p].update(nodes[u], inf)
		heap.Pop(&sons[p])
		nodes[u].value = newValue
		sum[u] -= sum[v]
		heap.Push(&sons[v], nodes[u])
		fa[u] = v
		fa[v] = p
		sum[v] += sum[u]
		nodes[v].value += nodes[u].value
		heap.Push(&sons[p], nodes[v])
	}

	var ans []int

	for _, cur := range Q {
		if cur[0] == 1 {
			x := cur[1] - 1
			ans = append(ans, sum[x])
		} else {
			update(cur[1] - 1)
		}
	}

	return ans
}

type Node struct {
	id    int
	value int
	index int
}

type PQ []*Node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].value > pq[j].value || pq[i].value == pq[j].value && pq[i].id < pq[j].id
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := x.(*Node)
	n.index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

func (pq *PQ) update(node *Node, value int) {
	node.value = value
	heap.Fix(pq, node.index)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
