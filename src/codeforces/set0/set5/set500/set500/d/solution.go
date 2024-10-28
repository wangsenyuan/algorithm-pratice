package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%.8f\n", x))
	}

	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) []float64 {
	n := readNum(reader)
	roads := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		roads[i] = readNNums(reader, 3)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, roads, queries)
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
func nC1(n int) int {
	return n
}

func nC2(n int) int {
	return n * (n - 1) / 2
}

func nC3(n int) int {
	return n * (n - 1) * (n - 2) / 6
}

func solve(n int, roads [][]int, queries [][]int) []float64 {
	g := NewGraph(n, len(roads)*2)

	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	sz := make([]int, n)
	fa := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				fa[v] = u
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	// 貌似会溢出呐
	base := nC3(n)

	var tot float64
	val := make([]int, len(roads))
	fac := make([]float64, len(roads))
	for i, road := range roads {
		u, v, w := road[0], road[1], road[2]
		u--
		v--
		if fa[u] == v {
			v = u
		}
		val[i] = w
		fac[i] = float64(nC1(sz[v])*nC2(n-sz[v]))/float64(base) + float64(nC2(sz[v])*nC1(n-sz[v]))/float64(base)
		tot += float64(val[i]) * fac[i]
	}

	ans := make([]float64, len(queries))

	for i, cur := range queries {
		eid, x := cur[0], cur[1]
		eid--
		w := val[eid]

		tot -= float64(w-x) * fac[eid]

		ans[i] = float64(2 * tot)

		val[eid] = x
	}

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.val = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
