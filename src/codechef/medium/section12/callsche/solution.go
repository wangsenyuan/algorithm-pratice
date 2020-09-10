package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		person_count := line[0]
		days_in_week := line[1]
		working_day_center := line[2]
		working_day_person := line[3]
		L := readNNums(reader, person_count)
		lunch_begin, lunch_end := readTwoNums(reader)
		R := make([][]int, person_count)
		for i := 0; i < person_count; i++ {
			R[i] = readNNums(reader, working_day_center)
		}

		F := make([][][]int, person_count)
		for i := 0; i < person_count; i++ {
			F[i] = make([][]int, days_in_week)
			for j := 0; j < days_in_week; j++ {
				F[i][j] = readNNums(reader, working_day_center)
			}
		}
		res := solve(person_count, days_in_week, working_day_center, working_day_person, L, lunch_begin, lunch_end, R, F)

		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func solve(people_count, days_in_week, working_day_center, working_day_person int,
	L []int, lunch_begin, lunch_end int, R [][]int, F [][][]int) bool {
	var nodes int
	source := nodes
	nodes++
	sink := nodes
	nodes++
	person := make([]int, 0, people_count)
	for i := 0; i < people_count; i++ {
		person = append(person, nodes)
		nodes++
	}
	person_day := make([][]int, people_count)
	for i := 0; i < people_count; i++ {
		person_day[i] = make([]int, 0, days_in_week)
		for j := 0; j < days_in_week; j++ {
			person_day[i] = append(person_day[i], nodes)
			nodes++
		}
	}

	person_day_lunch := make([][]int, people_count)
	for i := 0; i < people_count; i++ {
		person_day_lunch[i] = make([]int, 0, days_in_week)
		for j := 0; j < days_in_week; j++ {
			person_day_lunch[i] = append(person_day_lunch[i], nodes)
			nodes++
		}
	}

	day_hour := make([][]int, days_in_week)
	for i := 0; i < days_in_week; i++ {
		day_hour[i] = make([]int, 0, working_day_center)
		for j := 0; j < working_day_center; j++ {
			day_hour[i] = append(day_hour[i], nodes)
			nodes++
		}
	}

	lunch_begin--
	lunch_end--

	var total int
	for i := 0; i < len(R); i++ {
		for j := 0; j < len(R[i]); j++ {
			total += R[i][j]
		}
	}

	graph := NewFlowGraph(nodes, source, sink)

	var good = true

	for i := 0; i < people_count; i++ {
		graph.AddEdge(source, person[i], L[i])
		for j := 0; j < days_in_week; j++ {
			meetings_free := working_day_person
			for k := 0; k < working_day_center; k++ {
				meetings_free -= 1 ^ F[i][j][k]
			}
			if meetings_free < 0 {
				good = false
			}
			graph.AddEdge(person[i], person_day[i][j], meetings_free)
			var lunch_hour_free int
			for k := lunch_begin; k <= lunch_end; k++ {
				lunch_hour_free += F[i][j][k]
			}
			if lunch_hour_free == 0 {
				good = false
			}
			graph.AddEdge(person[i], person_day_lunch[i][j], lunch_hour_free-1)

			for k := 0; k < working_day_center; k++ {
				from := person_day_lunch[i][j]
				if k < lunch_begin || k > lunch_end {
					from = person_day[i][j]
				}
				graph.AddEdge(from, day_hour[j][k], F[i][j][k])
			}
		}
	}

	for i := 0; i < days_in_week; i++ {
		for j := 0; j < working_day_center; j++ {
			graph.AddEdge(day_hour[i][j], sink, R[i][j])
		}
	}

	return good && graph.MaxFlow() == total
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
