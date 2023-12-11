package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	P := readNNums(reader, n)
	m := readNum(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}

	res := solve(n, P, E)

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

func solve(n int, P []int, E [][]int) int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = -1
		}
	}

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		if mat[u][v] < 0 || mat[u][v] > w {
			mat[u][v] = w
		}
	}

	g := NewGraph(n, n*n)

	verts := make([]*Vertex, n)
	pq := make(PQ, n)
	root := -1

	for i := 0; i < n; i++ {
		v := new(Vertex)
		v.id = i
		v.cost = inf
		v.index = i
		verts[i] = v
		pq[i] = v

		if root < 0 || P[root] < P[i] {
			root = i
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || mat[i][j] < 0 {
				continue
			}
			g.AddEdge(i, j, mat[i][j])
			verts[j].deg++
		}
	}

	for i := 0; i < n; i++ {
		if root != i && P[root] == P[i] {
			return -1
		}
		if root != i && verts[i].deg == 0 {
			// multiple roots
			return -1
		}
	}

	// P[ai] > P[aj] => 这个是一个dag， 所以只能从最大值开始
	// 有两个最大值，就没有答案，然后对于任何一个节点，最好是选择最小的cost
	// 从某个节点开始，构建出一棵树
	heap.Init(&pq)

	pq.update(verts[root], 0, 0)

	var res int

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Vertex)

		u := cur.id

		if cur.deg > 0 || cur.cost == inf {
			return -1
		}
		res += cur.cost

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if verts[v].index < 0 {
				continue
			}
			// any way it will decrease
			pq.update(verts[v], verts[v].deg-1, min(verts[v].cost, g.val[i]))
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

type Vertex struct {
	id    int
	deg   int
	cost  int
	index int
}

type PQ []*Vertex

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].deg < pq[j].deg || pq[i].deg == pq[j].deg && pq[i].cost < pq[j].cost
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	v := x.(*Vertex)
	v.index = len(*pq)
	*pq = append(*pq, v)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PQ) update(v *Vertex, deg int, cost int) {
	v.cost = cost
	v.deg = deg
	heap.Fix(pq, v.index)
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
