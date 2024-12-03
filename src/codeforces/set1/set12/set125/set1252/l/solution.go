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

	res, _, _, _ := process(reader)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func process(reader *bufio.Reader) (res [][]int, a []int, cities [][]int, workers []int) {
	n, k := readTwoNums(reader)
	a = make([]int, n)
	cities = make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &a[i]) + 1
		var m int
		pos = readInt(s, pos, &m) + 1
		cities[i] = make([]int, m)
		for j := range m {
			pos = readInt(s, pos, &cities[i][j]) + 1
		}
	}
	workers = readNNums(reader, k)
	res = solve(a, cities, workers)
	return
}

func solve(a []int, cities [][]int, workers []int) [][]int {
	n := len(cities)
	m := len(workers)
	// a 似乎没有用
	materials := make(map[int]int)

	for _, x := range workers {
		materials[x]++
	}
	arr := make([]int, 0, len(materials))
	for k := range materials {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	for i, k := range arr {
		materials[k] = i
	}
	k := len(arr)

	cnt := make([]int, k)

	for _, x := range workers {
		j := materials[x]
		cnt[j]++
	}

	deg := make([]int, n)
	for i := 0; i < n; i++ {
		a[i]--
		deg[a[i]]++
	}

	marked := make([]bool, n)

	que := make([]int, n)
	var head, tail int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[head] = i
			head++
			marked[i] = true
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		deg[a[u]]--
		if deg[a[u]] == 0 {
			que[head] = a[u]
			head++
			marked[a[u]] = true
		}
	}

	source := n + k
	sink := source + 1
	cycle_point := sink + 1

	g := NewFlowGraph(cycle_point+1, source, sink)

	g.AddEdge(source, cycle_point, n-head-1)

	for i := 0; i < n; i++ {
		if !marked[i] && !marked[a[i]] {
			g.AddEdge(cycle_point, i, 1)
		} else {
			// i 不在cycle上
			g.AddEdge(source, i, 1)
		}
		for _, x := range cities[i] {
			if j, ok := materials[x]; ok {
				g.AddEdge(i, n+j, 1)
			}
		}
	}

	for i := 0; i < k; i++ {
		// 这个材料有这么多人可以工作，所以它的容量就是worker的数量
		g.AddEdge(n+i, sink, cnt[i])
	}

	res := g.MaxFlow()

	if res < n-1 {
		return nil
	}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = []int{0, 0}
		j := materials[workers[i]] + n

		for _, eid := range g.adj[j] {
			e := g.edges[eid]
			if e.flow == -1 && e.to < n {
				u := e.to
				ans[i][0] = u + 1
				ans[i][1] = a[u] + 1
				e.flow = 0
				break
			}
		}
	}

	return ans
}

const INF = 1000000000

type FlowGraph struct {
	n            int
	edges        []*Edge
	adj          [][]int
	que          []int
	dist         []int
	source, sink int
}

func NewFlowGraph(n int, source int, sink int) FlowGraph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, 10)
	}
	edges := make([]*Edge, 0, n)
	return FlowGraph{n, edges,
		adj, make([]int, n), make([]int, n), source, sink}
}

func (g *FlowGraph) AddEdge(u, v, c int) {
	g.adj[u] = append(g.adj[u], len(g.edges))
	e1 := &Edge{u, v, c, 0}
	g.edges = append(g.edges, e1)

	g.adj[v] = append(g.adj[v], len(g.edges))
	e2 := &Edge{v, u, 0, 0}
	g.edges = append(g.edges, e2)
}

func (g FlowGraph) bfs() bool {
	for i := 0; i < g.n; i++ {
		g.dist[i] = -1
	}
	var head, tail int

	g.que[tail] = g.source
	tail++
	g.dist[g.source] = 0

	for head < tail {
		u := g.que[head]
		head++
		for _, id := range g.adj[u] {
			e := g.edges[id]
			if g.dist[e.to] < 0 && e.flow < e.capacity {
				g.que[tail] = e.to
				tail++
				g.dist[e.to] = g.dist[u] + 1
			}
		}
	}
	return g.dist[g.sink] != -1
}

func (g *FlowGraph) dfs(v int, flow int, ptr []int) int {
	if flow == 0 {
		return 0
	}
	if v == g.sink {
		return flow
	}

	for ptr[v] < len(g.adj[v]) {
		id := g.adj[v][ptr[v]]
		ptr[v]++
		edge := g.edges[id]
		if g.dist[edge.to] != g.dist[edge.from]+1 {
			continue
		}
		pushed := g.dfs(edge.to, min(flow, edge.capacity-edge.flow), ptr)

		if pushed > 0 {
			g.edges[id].flow += pushed
			g.edges[id^1].flow -= pushed
			return pushed
		}
	}
	return 0
}

func (g *FlowGraph) MaxFlow() int {
	var flow int
	ptr := make([]int, g.n)
	for g.bfs() {
		for i := 0; i < g.n; i++ {
			ptr[i] = 0
		}
		for {
			pushed := g.dfs(g.source, INF, ptr)
			if pushed == 0 {
				break
			}
			flow += pushed
		}
	}

	return flow
}

type Edge struct {
	from, to       int
	capacity, flow int
}
