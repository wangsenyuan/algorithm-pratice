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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readString(reader)
		n, m := readTwoNums(reader)
		s, t := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, s, t)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}
func solve(n int, E [][]int, s int, t int) int {
	s--
	t--
	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)
	var front, end int
	enqueue := func(u int) {
		que[end] = u
		end++
	}
	dequeue := func() int {
		u := que[front]
		front++
		return u
	}
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	dist[s] = 0

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[s][0] = 1

	enqueue(s)
	for front < end {
		u := dequeue()
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				enqueue(v)
				dp[v][0] = dp[u][0]
			} else if dist[v] == dist[u]+1 {
				dp[v][0] = add(dp[v][0], dp[u][0])
			}
		}
	}

	type Pair struct {
		first  int
		second int
	}

	xx := make([]Pair, n)
	for i := 0; i < n; i++ {
		xx[i] = Pair{dist[i], i}
	}

	sort.Slice(xx, func(i, j int) bool {
		return xx[i].first < xx[j].first
	})

	for i := 1; i < n; {
		j := i
		for ; i < n && xx[i].first == xx[j].first; i++ {
			u := xx[i].second
			for k := g.nodes[u]; k > 0; k = g.next[k] {
				v := g.to[k]
				if dist[v] == dist[u] {
					dp[u][1] = add(dp[u][1], dp[v][0])
				} else if dist[v]+1 == dist[u] {
					dp[u][1] = add(dp[u][1], dp[v][1])
				}
			}
		}
	}

	return add(dp[t][0], dp[t][1])
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
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
