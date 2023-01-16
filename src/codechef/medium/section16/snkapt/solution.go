package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line := readNNums(reader, 4)
	n, m, z, t := line[0], line[1], line[2], line[3]
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}
	trees := make([][]int, z)
	for i := 0; i < z; i++ {
		trees[i] = readNNums(reader, 5)
	}
	res := solve(t, mat, trees)
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

var dd []int = []int{-1, 0, 1, 0, -1}

func solve(T int, field []string, trees [][]int) int {
	var snakes []Pair

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == 'S' {
				snakes = append(snakes, Pair{i, j})
			}
		}
	}

	var ctr int
	var inv []Triple

	inv = append(inv, Triple{-1, -1, -1})

	indexes := make([][][]int, len(field))

	for i := 0; i < len(field); i++ {
		indexes[i] = make([][]int, len(field[i]))
		for j := 0; j < len(field[i]); j++ {
			indexes[i][j] = make([]int, T+1)
			if field[i][j] != '#' {
				for k := 0; k <= T; k++ {
					ctr += 2
					indexes[i][j][k] = ctr
					inv = append(inv, Triple{i, j, k})
				}
			}
		}
	}
	ctr++
	last := ctr
	ctr++

	cost := make([][][]int, len(field))
	for i := 0; i < len(field); i++ {
		cost[i] = make([][]int, len(field[i]))
		for j := 0; j < len(field[i]); j++ {
			cost[i][j] = make([]int, T+1)
		}
	}

	for _, tree := range trees {
		xi, yi, pi, qi, hi := tree[0], tree[1], tree[2], tree[3], tree[4]
		xi--
		yi--

		for t := pi; t < qi; t++ {
			cost[xi][yi][t] = min(cost[xi][yi][t], -hi)
		}
	}

	countOrCreate := func(g *Graph) int {
		var cnt int
		for x := 0; x < len(field); x++ {
			for y := 0; y < len(field[x]); y++ {
				if field[x][y] != '#' {
					for j := 0; j < 4; j++ {
						u := x + dd[j]
						v := y + dd[j+1]
						if u >= 0 && u < len(field) && v >= 0 && v < len(field[x]) && field[u][v] != '#' {
							for k := 0; k < T; k++ {
								cnt++
								if g != nil {
									a := indexes[x][y][k]
									b := indexes[u][v][k+1] - 1
									g.AddEdge(a, b, 1, 0)
								}
							}
						}
					}

					for k := 0; k < T; k++ {
						cnt++
						if g != nil {
							a := indexes[x][y][k]
							b := indexes[x][y][k+1] - 1
							g.AddEdge(a, b, 1, cost[x][y][k])
						}
					}
				}
			}
		}

		for i := 1; i < last; i += 2 {
			cnt++
			if g != nil {
				g.AddEdge(i, i+1, 1, 0)
			}
		}
		for i := 1; i < len(inv); i++ {
			if inv[i].k == T {
				cnt++
				if g != nil {
					g.AddEdge(2*i, last, 1, 0)
				}
			}
		}

		for _, x := range snakes {
			cnt++
			if g != nil {
				a := indexes[x.first][x.second][0] - 1
				g.AddEdge(0, a, 1, 0)
			}
		}

		return cnt
	}
	// g := NewGraph(ctr, 10*ctr)
	edgeCount := countOrCreate(nil)
	g := NewGraph(ctr, 2*edgeCount)
	countOrCreate(g)

	res := minCostFlow(g, 0, last, len(snakes))

	return -res
}

type Triple struct {
	i int
	j int
	k int
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	flow  []int
	cap   []int
	cost  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.flow = make([]int, e+2)
	g.cap = make([]int, e+2)
	g.cost = make([]int, e+2)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, cap, cost int) {

	add := func(x, y, cap, cost int) {
		g.cur++
		g.next[g.cur] = g.nodes[x]
		g.nodes[x] = g.cur
		g.to[g.cur] = y
		g.cap[g.cur] = cap
		g.cost[g.cur] = cost
		g.flow[g.cur] = 0
	}
	add(u, v, cap, cost)
	add(v, u, 0, -cost)
}

func bellmanFord(g *Graph, s int) []int {
	n := len(g.nodes)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = 1 << 30
	}
	dist[s] = 0
	que := make([]int, n)
	var front, end int
	que[end] = s
	end++
	in := make([]bool, n)
	for (end-front)%n != 0 {
		u := que[front%n]
		front++
		// front = (front + 1) % n
		in[u] = false
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			if g.cap[i] > g.flow[i] {
				// direct edge
				v := g.to[i]
				if dist[v] > dist[u]+g.cost[i] {
					dist[v] = dist[u] + g.cost[i]
					if !in[v] {
						in[v] = true
						que[end%n] = v
						end++
					}
				}
			}
		}
	}

	return dist
}

type Pair struct {
	first  int
	second int
}

type PQ []Pair

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

func minCostFlow(g *Graph, src int, dst int, maxflow int) int {
	pot := bellmanFord(g, src)

	var flow, flowCost int

	n := len(g.nodes)
	pq := make(PQ, 0, n)

	prio := make([]int, n)
	curflow := make([]int, n)

	in := make([]Pair, n)

	for flow < maxflow {

		for i := 0; i < n; i++ {
			prio[i] = 1 << 30
		}
		heap.Push(&pq, Pair{0, src})
		curflow[src] = 1 << 30
		prio[src] = 0

		for pq.Len() > 0 {
			p := heap.Pop(&pq).(Pair)
			u := p.second
			if p.first != prio[u] {
				continue
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				if g.flow[i] >= g.cap[i] {
					continue
				}
				v := g.to[i]
				newPrio := prio[u] + g.cost[i] + pot[u] - pot[v]
				if prio[v] > newPrio {
					prio[v] = newPrio
					heap.Push(&pq, Pair{prio[v], v})
					curflow[v] = min(curflow[u], g.cap[i]-g.flow[i])
					in[v] = Pair{u, i}
				}
			}
		}

		if prio[dst] == 1<<30 {
			break
		}
		for i := 0; i < n; i++ {
			pot[i] += prio[i]
		}
		df := min(curflow[dst], maxflow-flow)
		flow += df

		for v := dst; v != src; v = in[v].first {
			i := in[v].second
			g.flow[i] += df
			g.flow[i^1] -= df
			flowCost += df * g.cost[i]
		}
	}

	return flowCost
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
