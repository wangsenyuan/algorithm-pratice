package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		p := readNNums(reader, n-1)
		res := solve(n, a, p)

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

const inf = 1 << 50

func solve(n int, a []int, p []int) int {
	if n == 1 {
		return a[0]
	}

	g := NewGraph(n, n)

	for i := 0; i < n-1; i++ {
		g.AddEdge(p[i]-1, i+1, -1)
	}

	var dfs func(u int) int

	dfs = func(u int) int {
		if g.nodes[u] == 0 {
			return a[u]
		}
		res := inf
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			res = min(res, dfs(v))
		}

		return min(res, (res+a[u])/2)
	}

	res := inf
	for i := g.nodes[0]; i > 0; i = g.next[i] {
		v := g.to[i]
		res = min(res, dfs(v))
	}

	return a[0] + res
}

func solve1(n int, a []int, p []int) int {
	g := NewGraph(n, n)

	for i := 0; i < n-1; i++ {
		g.AddEdge(p[i]-1, i+1, -1)
	}
	ma := slices.Max(a)

	var dfs func(u int, sub int) bool

	dfs = func(u int, sub int) bool {
		if sub > ma {
			return false
		}
		cur := a[u] - sub

		if g.nodes[u] == 0 {
			return cur >= 0
		}

		if cur < 0 {
			sub -= cur
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, sub) {
				return false
			}
		}
		return true
	}

	check := func(expect int) bool {
		sub := expect - a[0]
		for i := g.nodes[0]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, sub) {
				return false
			}
		}
		return true
	}

	l, r := a[0], a[0]+ma+1

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r - 1
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
