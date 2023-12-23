package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	P := readNNums(reader, n-1)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, P, A, Q)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(n int, P []int, A []int, Q [][]int) []int {
	// x, y
	g := NewGraph(n, n)

	for i := 1; i < n; i++ {
		p := P[i-1] - 1
		g.AddEdge(p, i)
	}

	depth := make([]int, n)
	cnt := make([]int, n)
	var dfs func(u int)
	dfs = func(u int) {
		cnt[depth[u]]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			depth[v] = depth[u] + 1
			dfs(v)
		}
	}

	dfs(0)

	pos := make([]int, n)
	fp := make([][]int, n)

	sn := int(math.Sqrt(float64(n))) + 10

	for i := 0; i < n; i++ {
		fp[i] = make([]int, min(sn, cnt[depth[i]]))
		for j := 0; j < len(fp[i]); j++ {
			fp[i][j] = -1
		}
	}

	var dfs2 func(u int)

	dfs2 = func(u int) {
		d := depth[u]
		cnt[d]--
		pos[u] = cnt[d]

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs2(v)
		}
	}

	dfs2(0)

	var f func(x int, y int) int
	f = func(x int, y int) int {
		if x == 0 {
			return A[0] * A[0]
		}
		if pos[y] < len(fp[x]) && fp[x][pos[y]] > 0 {
			return fp[x][pos[y]]
		}

		res := f(P[x-1]-1, P[y-1]-1) + A[x]*A[y]

		if pos[y] < len(fp[x]) {
			fp[x][pos[y]] = res
		}

		return res
	}

	ans := make([]int, len(Q))

	for i, cur := range Q {
		x, y := cur[0], cur[1]
		x--
		y--
		ans[i] = f(x, y)
	}

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
