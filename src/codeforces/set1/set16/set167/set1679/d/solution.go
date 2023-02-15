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

	tc := 1

	for tc > 0 {
		tc--
		var n, m int
		var k int64
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &n)
		pos = readInt(s, pos+1, &m)
		readInt64(s, pos+1, &k)
		A := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, A, E, k)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1 << 30

func solve(n int, A []int, E [][]int, k int64) int {
	g := NewGraph(n, len(E))
	P := make([]Pair, n)

	for i := 0; i < n; i++ {
		P[i] = Pair{A[i], i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	que := make([]int, n)
	deg := make([]int, n)
	dist := make([]int, n)

	check := func(x int) bool {
		g.Reset()
		for i := 0; i < n; i++ {
			deg[i] = 0
			dist[i] = 0
		}
		for _, e := range E {
			u, v := e[0]-1, e[1]-1
			if A[u] <= x && A[v] <= x {
				g.AddEdge(u, v)
				deg[v]++
			}
		}
		var front, end int
		for u := 0; u < n; u++ {
			if deg[u] == 0 {
				que[end] = u
				end++
			}
		}
		var tot int
		for front < end {
			u := que[front]
			front++
			tot++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				deg[v]--
				dist[v] = max(dist[v], dist[u]+1)
				if deg[v] == 0 {
					que[end] = v
					end++
				}
			}
		}

		if tot < n {
			return true
		}

		for u := 0; u < n; u++ {
			if int64(dist[u]) >= k-1 {
				return true
			}
		}

		return false
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if check(P[mid].first) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r == n {
		return -1
	}
	return P[r].first
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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

func (g *Graph) Reset() {
	g.cur = 0
	for i := 0; i < len(g.nodes); i++ {
		g.nodes[i] = 0
	}
}
