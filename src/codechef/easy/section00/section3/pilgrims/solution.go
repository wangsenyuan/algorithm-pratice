package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([]int64, m)
		var pos int
		s, _ := reader.ReadBytes('\n')
		for i := 0; i < m; i++ {
			pos = readInt64(s, pos, &E[i]) + 1
		}
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 3)
		}
		res := solve(n, edges, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * int64(sign)
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	var tmp int64
	from = readInt64(bytes, from, &tmp)
	*val = int(tmp)
	return from
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

func solve(n int, edges [][]int, E []int64) int {
	g := NewGraph(n, len(edges)*2)

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	var dfs func(u, d int)
	leaves := make([]int, 0, n)
	F := make([]int64, n)
	for i := 0; i < n; i++ {
		F[i] = -1
	}
	dfs = func(u, d int) {
		var leaf = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if F[v] < 0 {
				F[v] = F[u] + int64(d)*int64(g.val[i])
				dfs(v, d+1)
				leaf = false
			}
		}

		if leaf {
			leaves = append(leaves, u)
		}
	}

	F[0] = 0

	dfs(0, 1)

	L := make([]int64, len(leaves))

	for i := 0; i < len(leaves); i++ {
		L[i] = F[leaves[i]]
	}

	sort.Sort(Longs(E))
	sort.Sort(Longs(L))

	var res int

	for i, j := 0, 0; i < len(L); i++ {
		for j < len(E) && E[j] < L[i] {
			j++
		}
		if j < len(E) {
			res++
			j++
		}
	}
	return res
}

type Longs []int64

func (this Longs) Len() int {
	return len(this)
}

func (this Longs) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Longs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int, e+10)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
