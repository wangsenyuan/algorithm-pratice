package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)

		X := readNNums(reader, n)

		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, X)

		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, E [][]int, X []int) []int {

	g := buildTree(n, E)

	cnt := make([][]int, n)

	dfs(-1, 0, cnt, X, g)

	diff := cnt[0][1] - cnt[0][0]

	if diff == 0 {
		return []int{1}
	}

	var res []int

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		if diff == 0 {
			return
		}
		cur := 2 * (cnt[u][1] - cnt[u][0])

		if diff > 0 && cur > 0 && cur <= diff {
			res = append(res, u+1)
			diff -= cur
			return
		}

		if diff < 0 && cur < 0 && cur >= diff {
			res = append(res, u+1)
			diff -= cur
			return
		}

		// cur > diff
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
			}
		}
	}

	dfs2(-1, 0)

	return res
}

func dfs(p, u int, cnt [][]int, X []int, g *Graph) {
	cnt[u] = make([]int, 2)
	for i := g.node[u]; i > 0; i = g.next[i] {
		v := g.to[i]
		if p == v {
			continue
		}
		dfs(u, v, cnt, X, g)
		for j := 0; j < 2; j++ {
			cnt[u][j] += cnt[v][j]
		}
	}
	cnt[u][X[u]]++
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func buildTree(n int, E [][]int) *Graph {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	return g
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
