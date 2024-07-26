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
	n, m, k := readThreeNums(reader)
	controls := make([][]int, n)
	for i := 0; i < n; i++ {
		controls[i] = readNNums(reader, 2)
	}
	fences := make([][]int, m)
	for i := 0; i < m; i++ {
		fences[i] = readNNums(reader, 3)
	}
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(controls, fences, queries)

	var buf bytes.Buffer
	for i := 0; i < k; i++ {
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

type cycle struct {
	x int
	y int
	r int
}

const inf = 1 << 32

const H = 11

func solve(controls [][]int, fences [][]int, queries [][]int) []int {
	m := len(fences)
	cycles := make([]cycle, m+1)
	cycles[0] = cycle{0, 0, 2 * inf}

	for i := 0; i < m; i++ {
		cycles[i+1] = cycle{fences[i][1], fences[i][2], fences[i][0]}
	}

	sort.Slice(cycles, func(i, j int) bool {
		return cycles[i].r < cycles[j].r
	})

	g := NewGraph(m+1, m+1)
	for i := 0; i < m; i++ {
		for j := i + 1; j <= m; j++ {
			// 因为两个fence没有重叠的地方，
			// 如果 dx * dx + dy * dy < (r1 + r2) * (r1 + r2)
			dx := cycles[i].x - cycles[j].x
			dy := cycles[i].y - cycles[j].y
			r := cycles[i].r + cycles[j].r
			if dx*dx+dy*dy < r*r {
				g.AddEdge(j, i)
				break
			}
		}
	}

	depth := make([]int, m+1)
	fa := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		fa[i] = make([]int, H)
	}
	for i := 0; i < H; i++ {
		fa[m][i] = m
	}
	var dfs func(u int)
	dfs = func(u int) {
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			fa[v][0] = u
			depth[v] = depth[u] + 1
			dfs(v)
		}
	}

	dfs(m)

	ans := make([]int, len(queries))

	findCycle := func(ctr []int) int {
		x, y := ctr[0], ctr[1]
		for i := 0; i < m; i++ {
			cur := cycles[i]
			dx := x - cur.x
			dy := y - cur.y
			r := cur.r
			if dx*dx+dy*dy < r*r {
				return i
			}
		}

		return m
	}

	lca := func(u int, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if depth[u]-1<<i >= depth[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	for i, cur := range queries {
		a, b := cur[0]-1, cur[1]-1
		u := findCycle(controls[a])
		v := findCycle(controls[b])
		p := lca(u, v)
		ans[i] = depth[u] + depth[v] - 2*depth[p]
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
