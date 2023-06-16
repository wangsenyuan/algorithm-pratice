package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)
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

const H = 18

func solve(n int, E [][]int) int64 {
	g := NewGraph(n, len(E)*2)
	deg := make([]int, n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int)

	fa := make([][]int, n)
	depth := make([]int, n)
	dfs = func(p int, u int) {
		fa[u] = make([]int, H)
		fa[u][0] = p
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				depth[v] = depth[u] + 1
				dfs(u, v)
			}
		}
	}

	depth[0] = 1

	dfs(0, 0)

	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]int, H)
		pos[i][0] = deg[i]
	}

	its := make([]Item, n)

	for step := 1; step < H; step++ {
		for i := 0; i < n; i++ {
			first := pos[i][step-1]
			second := -1
			if depth[i] > (1 << (step - 1)) {
				second = pos[fa[i][step-1]][step-1]
			}
			its[i] = Item{first, second, i}
		}
		sort.Slice(its, func(i, j int) bool {
			return its[i].first < its[j].first || its[i].first == its[j].first && its[i].second < its[j].second
		})

		pos[its[0].id][step] = 0
		for i := 1; i < n; i++ {
			pos[its[i].id][step] = pos[its[i-1].id][step]
			if its[i].first != its[i-1].first || its[i].second != its[i-1].second {
				pos[its[i].id][step]++
			}
		}
	}

	lcp := func(u, v int) int {
		if pos[u][H-1] == pos[v][H-1] {
			return min(depth[u], depth[v])
		}
		var ln int
		for k := H - 1; k >= 0; k-- {
			if pos[u][k] == pos[v][k] {
				ln += 1 << k
				u = fa[u][k]
				v = fa[v][k]
			}
		}
		return ln
	}

	res := int64(depth[its[0].id])
	for i := 1; i < n; i++ {
		res += int64(depth[its[i].id])
		res -= int64(lcp(its[i].id, its[i-1].id))
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	first  int
	second int
	id     int
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
	return &Graph{nodes, next, to, -1}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
