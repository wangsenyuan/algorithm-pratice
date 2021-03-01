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
		n, m := readTwoNums(reader)
		P := readNNums(reader, n-1)
		A := make([][]int, m)
		for i := 0; i < m; i++ {
			A[i] = readByFirst(reader)
		}
		res := solve(n, m, P, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readByFirst(reader *bufio.Reader) []int {
	s, _ := reader.ReadBytes('\n')
	var n int
	pos := readInt(s, 0, &n) + 1
	res := make([]int, n)
	for i := 0; i < n; i++ {
		pos = readInt(s, pos, &res[i]) + 1
	}
	return res
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

type Graph struct {
	nodes  []int
	next   []int
	to     []int
	weight []int
	cur    int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	weight := make([]int, e+10)
	cur := 1
	return &Graph{nodes, next, to, weight, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.weight[g.cur] = w
}

const INF = 1 << 30

func solve(n, m int, P []int, A [][]int) int {
	// 0 is source, and n + 1 is sink, m is for L part
	N := n + 2 + m
	// n - 1 is for tree-parent child, next m is for source to nodes, and last n for sink
	E := 2*(n-1) + 2*m + 2*n
	for _, a := range A {
		// L to R
		E += len(a) * 2
	}
	g := NewGraph(N, E)

	for i := 2; i <= n; i++ {
		p := P[i-2]
		// from parent to child has inf capacity
		g.AddEdge(p, i, INF)
		// from child to parent has 0 capacity
		g.AddEdge(i, p, 0)
	}

	// add source to L edges
	for i := 1; i <= m; i++ {
		g.AddEdge(0, n+i, 1)
		g.AddEdge(n+i, 0, 0)
	}

	for i := 1; i <= m; i++ {
		cur := A[i-1]
		for _, j := range cur {
			g.AddEdge(n+i, j, 1)
			g.AddEdge(j, n+i, 0)
		}
	}

	for i := 1; i <= n; i++ {
		g.AddEdge(i, N-1, 1)
		g.AddEdge(N-1, i, 0)
	}

	que := make([]int, N)
	level := make([]int, N)
	bfs := func(s int, t int) bool {
		for i := 0; i < N; i++ {
			level[i] = -1
		}
		var front, end int
		que[end] = s
		end++
		level[s] = 0
		for front < end {
			u := que[front]
			front++

			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if level[v] < 0 && g.weight[i] > 0 {
					level[v] = level[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return level[t] > 0
	}

	pos := make([]int, N)

	var send func(flow int, u, t int) int

	send = func(flow int, u, t int) int {
		if u == t {
			return flow
		}

		for ; pos[u] > 0; pos[u] = g.next[pos[u]] {
			i := pos[u]
			v := g.to[i]
			if level[v] == level[u]+1 && g.weight[i] > 0 {
				curFlow := min(flow, g.weight[i])
				tmp := send(curFlow, v, t)
				if tmp > 0 {
					g.weight[i] -= tmp
					g.weight[i^1] += tmp
					return tmp
				}
			}
		}
		return 0
	}

	var res int

	for bfs(0, N-1) {
		for i := 0; i < N; i++ {
			pos[i] = g.nodes[i]
		}
		for {
			flow := send(INF, 0, N-1)
			if flow == 0 {
				break
			}
			res += flow
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
