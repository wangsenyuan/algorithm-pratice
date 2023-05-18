package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	//var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		B := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, k, E, B)
		//buf.WriteString(fmt.Sprintf("%d\n", res))
		fmt.Println(res)
	}

	//fmt.Print(buf.String())
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const MOD = 1000000021

func solve(n, K int, E [][]int, B []int) int64 {
	g := NewGraph(n, len(E))
	rg := NewGraph(n, len(E))
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		rg.AddEdge(v, u)
	}
	vis := make([]bool, n)
	var stack []int

	var dfs func(u int, g *Graph)
	dfs = func(u int, g *Graph) {
		vis[u] = true
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v, g)
			}
		}
		stack = append(stack, u)
	}

	for u := 0; u < n; u++ {
		if !vis[u] {
			dfs(u, g)
		}
	}

	comp := make([]int, n)

	for i := 0; i < n; i++ {
		vis[i] = false
	}

	var cur []int64

	var dfs1 func(u int, gid int)

	dfs1 = func(u int, gid int) {
		cur = append(cur, int64(B[u]))
		comp[u] = gid
		vis[u] = true
		for i := rg.node[u]; i > 0; i = rg.next[i] {
			v := rg.to[i]
			if !vis[v] {
				dfs1(v, gid)
			}
		}
	}

	var fp [][]int64
	for i := n - 1; i >= 0; i-- {
		u := stack[i]
		if !vis[u] {
			cur = make([]int64, 0, 1)
			dfs1(u, len(fp))
			sort.Slice(cur, func(a, b int) bool {
				return cur[a] < cur[b]
			})
			fp = append(fp, cur)
		}
	}

	m := len(fp)
	ng := NewGraph(m, len(E))

	sets := make([]map[int]bool, m)
	for i := 0; i < m; i++ {
		sets[i] = make(map[int]bool)
	}

	for _, e := range E {
		u, v := comp[e[0]-1], comp[e[1]-1]
		if u != v && !sets[u][v] {
			ng.AddEdge(u, v)
			sets[u][v] = true
		}
	}
	stack = make([]int, 0, m)

	for i := 0; i < m; i++ {
		vis[i] = false
	}

	for i := 0; i < m; i++ {
		if !vis[i] {
			dfs(i, ng)
		}
	}

	var best int64

	tmp := make([]int64, K+1)

	dp := make([][]int64, m)

	for _, u := range stack {
		dp[u] = make([]int64, K+1)
		//u := stack[i]
		for j := ng.node[u]; j > 0; j = ng.next[j] {
			v := ng.to[j]
			for x := 0; x <= K; x++ {
				tmp[x] = max(tmp[x], dp[v][x])
			}
		}
		for j := 0; j <= K; j++ {
			for l := 0; l < len(fp[u]) && l <= j; l++ {
				dp[u][j] = max(dp[u][j], tmp[j-l]+int64(len(fp[u])-l)*fp[u][l])
			}
			best = max(best, dp[u][j])
		}

		for j := 0; j <= K; j++ {
			tmp[j] = 0
		}
	}

	return (best % MOD)
}

func max(a, b int64) int64 {
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
