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
		n, a, b := readThreeNums(reader)
		roads := make([][]int, n)
		for i := 0; i < n; i++ {
			roads[i] = readNNums(reader, 2)
		}
		res := solve(n, roads, a, b)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(n int, roads [][]int, a int, b int) bool {
	if a == b {
		return false
	}
	a--
	b--
	g := NewGraph(n, 2*n)

	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)

	trace := make([]int, n)
	var cycle []int

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if !vis[v] {
				trace[v] = u
				dfs(u, v)
			} else if len(cycle) == 0 {
				x := u
				for x != v {
					cycle = append(cycle, x)
					x = trace[x]
				}
				cycle = append(cycle, x)
			}
		}
	}

	dfs(-1, 0)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	que := make([]int, n)
	var front, tail int
	for _, u := range cycle {
		que[front] = u
		front++
		dist[u] = 0
		trace[u] = u
	}

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				trace[v] = trace[u]
				que[front] = v
				front++
			}
		}
	}

	if dist[b] == 0 {
		return true
	}
	if trace[a] == trace[b] {
		//Marcel can stop at the enter of cycle for Valeriu
		return dist[a] > dist[b]
	}

	critical := trace[b]

	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	dist[critical] = 0
	front = 0
	tail = 0
	que[front] = critical
	front++

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				trace[v] = trace[u]
				que[front] = v
				front++
			}
		}
	}

	return dist[a] > dist[b]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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
