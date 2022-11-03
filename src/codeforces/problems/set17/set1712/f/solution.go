package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	n := readNum(reader)
	P := readNNums(reader, n-1)
	m := readNum(reader)
	X := readNNums(reader, m)
	res := solve(n, P, X)
	for i := 0; i < m; i++ {
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

const INF = 1 << 30

func solve(n int, P []int, X []int) []int {
	// D(u, v) = min(fu + fv + x, d(u, v))
	g := NewGraph(n, 2*n)
	deg := make([]int, n)

	for i := 2; i <= n; i++ {
		p := P[i-2]
		deg[i-1]++
		deg[p-1]++
		g.AddEdge(i-1, p-1)
		g.AddEdge(p-1, i-1)
	}

	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = INF
	}

	que := make([]int, n)
	var front, end int
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			f[i] = 0
			que[end] = i
			end++
		}
	}

	for front < end {
		u := que[front]
		front++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if f[v] == INF {
				f[v] = f[u] + 1
				que[end] = v
				end++
			}
		}
	}

	depth := make([]int, n)

	val := make([][]int, n)
	for i := 0; i < n; i++ {
		val[i] = make([]int, 0, 1)
	}

	ans := make([]int, len(X))

	merge := func(x, y int) {
		if len(val[x]) < len(val[y]) {
			val[x], val[y] = val[y], val[x]
		}
		// merge y to x, and val[x] is what we want
		for id := 0; id < len(X); id++ {
			for i := 0; i < len(val[y]); i++ {
				j := max(0, ans[id]-i-X[id])
				for j < len(val[x]) && val[y][i]+val[x][j]-2*depth[x] >= ans[id] {
					ans[id]++
					j = max(0, ans[id]-i-X[id])
				}
			}
		}

		for i := 0; i < len(val[y]); i++ {
			val[x][i] = max(val[x][i], val[y][i])
		}
	}

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			depth[v] = depth[u] + 1
			dfs(u, v)
			merge(u, v)
		}

		for id := 0; id < len(X); id++ {
			j := max(0, ans[id]-f[u]-X[id])
			if j < len(val[u]) && val[u][j]+depth[u]-2*depth[u] >= ans[id] {
				ans[id]++
			}
		}

		if f[u] == len(val[u]) {
			val[u] = append(val[u], depth[u])
		}
	}

	dfs(0, 0)

	for i := 0; i < len(X); i++ {
		ans[i]--
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
	g.to = make([]int, e+1)
	g.next = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
