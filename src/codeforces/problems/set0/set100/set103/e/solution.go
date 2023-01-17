package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	sets := make([][]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadBytes('\n')
		var cnt int
		pos := readInt(line, 0, &cnt)
		sets[i] = make([]int, cnt)
		for j := 0; j < cnt; j++ {
			pos = readInt(line, pos+1, &sets[i][j])
		}
	}
	X := readNNums(reader, n)
	res := solve(n, sets, X)
	fmt.Println(res)
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const INF = 1e10

func solve(n int, sets [][]int, X []int) int64 {

	var edge int

	edge += 2 * n
	for _, set := range sets {
		edge += len(set)
	}
	var src int
	var dst int = 2*n + 1

	g := NewGraph(dst+1, 2*edge)

	for i := 1; i <= n; i++ {
		for _, x := range sets[i-1] {
			g.AddEdge(i, x+n, INF)
			g.AddEdge(x+n, i, 0)
		}
	}

	var res int64

	for i := 1; i <= n; i++ {
		res -= int64(X[i-1])
		g.AddEdge(src, i, INF-int64(X[i-1]))
		g.AddEdge(i, src, 0)
		g.AddEdge(i+n, dst, INF)
		g.AddEdge(dst, i+n, 0)
	}
	N := dst + 1
	que := make([]int, N)
	dis := make([]int, N)

	bfs := func() bool {
		for i := 0; i < N; i++ {
			dis[i] = -1
		}
		dis[src] = 0
		var front, end int
		que[end] = src
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dis[v] < 0 && g.cap[i] > 0 {
					dis[v] = dis[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return dis[dst] > 0
	}

	pos := make([]int, N)

	reset := func() {
		for i := 0; i < N; i++ {
			pos[i] = g.node[i]
		}
	}

	var dfs func(u int, mx int64) int64

	dfs = func(u int, mx int64) int64 {
		if u == dst {
			return mx
		}
		var sum int64
		for ; pos[u] > 0; pos[u] = g.next[pos[u]] {
			i := pos[u]
			v := g.to[i]
			c := g.cap[i]
			if dis[v] == dis[u]+1 && c > 0 {
				tmp := dfs(v, min(mx, c))
				if tmp == 0 {
					dis[v] = -1
				} else {
					mx -= tmp
					sum += tmp
					g.cap[i] -= tmp
					g.cap[i^1] += tmp
				}
			}
		}
		return sum
	}

	var ans int64

	for bfs() {
		reset()
		ans += dfs(src, INF)
	}

	res = res - (ans - int64(n)*INF)
	return -res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	cap  []int64
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cap = make([]int64, e+2)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v int, c int64) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.cap[g.cur] = c
}
