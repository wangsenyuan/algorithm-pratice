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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		h := readNNums(reader, n)
		deps := make([][]int, m)
		for i := 0; i < m; i++ {
			deps[i] = readNNums(reader, 2)
		}
		res := solve(k, h, deps)
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

func solve(k int, h []int, dependencies [][]int) int {
	n := len(h)

	g, deg := createGraph(n, dependencies)

	que := make([]int, n)

	var front, tail int

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[front] = i
			front++
		}
	}

	f := func(u int, v int) int {
		if h[u] <= h[v] {
			return 0
		}
		return 1
	}

	dp := make([]int, n)

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			dp[v] = max(dp[v], dp[u]+f(u, v)*k+h[v]-h[u])
			if deg[v] == 0 {
				que[front] = v
				front++
			}
		}
	}

	sort.Slice(que, func(i, j int) bool {
		return h[que[i]] < h[que[j]]
	})

	suf := make([]int, n+1)
	suf[n] = -(1 << 60)
	for i := n - 1; i >= 0; i-- {
		suf[i] = max(suf[i+1], dp[que[i]]-h[que[i]])
	}

	ans := 1 << 60
	pref := -ans

	for i := 0; i < n; i++ {
		j := que[i]
		pref = max(pref, dp[j]-h[j])
		ans = min(ans, max(pref+h[j], suf[i+1]+h[j]+k))
	}

	return ans
}

func createGraph(n int, dependencies [][]int) (*Graph, []int) {
	g := NewGraph(n, len(dependencies))
	deg := make([]int, n)
	for _, dep := range dependencies {
		a, b := dep[0], dep[1]
		a--
		b--
		g.AddEdge(a, b)
		deg[b]++
	}
	return g, deg
}

type Pair struct {
	first  int
	second int
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
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
