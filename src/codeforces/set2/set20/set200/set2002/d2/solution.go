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
		a := readNNums(reader, n-1)
		p := readNNums(reader, n)
		queries := make([][]int, m)
		for i := range m {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(n, a, p, queries)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
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

func solve(n int, a []int, p []int, queries [][]int) []bool {
	g := NewGraph(n, n)

	fa := make([]int, n)
	for i := 0; i < n-1; i++ {
		fa[i+1] = a[i] - 1
		g.AddEdge(fa[i+1], i+1)
	}

	in := make([]int, n)
	out := make([]int, n)
	var clock int

	var dfs func(u int)
	dfs = func(u int) {
		in[u] = clock
		clock++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			dfs(g.to[i])
		}
		out[u] = clock
	}

	dfs(0)

	for i := 0; i < n; i++ {
		p[i]--
	}

	check := func(i int) int {
		if p[i] == 0 {
			if i == 0 {
				return 1
			}
			return 0
		}
		if i == 0 {
			return 0
		}
		prev := p[i-1]

		if in[fa[p[i]]] <= in[prev] && out[prev] <= out[fa[p[i]]] {
			// 当前节点的父节点，是前一个节点的父节点
			return 1
		}
		return 0
	}
	var cnt int
	for i := 0; i < n; i++ {
		cnt += check(i)
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		x, y := cur[0]-1, cur[1]-1

		set := make(map[int]bool)
		set[x] = true
		set[y] = true
		if x-1 >= 0 {
			set[x-1] = true
		}
		if x+1 < n {
			set[x+1] = true
		}
		if y-1 >= 0 {
			set[y-1] = true
		}
		if y+1 < n {
			set[y+1] = true
		}

		for u := range set {
			cnt -= check(u)
		}
		p[x], p[y] = p[y], p[x]

		for u := range set {
			cnt += check(u)
		}

		ans[i] = cnt == n
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
