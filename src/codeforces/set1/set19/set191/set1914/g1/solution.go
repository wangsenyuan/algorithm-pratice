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
		n := readNum(reader)
		bulbs := readNNums(reader, 2*n)
		sz, ways := solve(n, bulbs)
		buf.WriteString(fmt.Sprintf("%d %d\n", sz, ways))
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, bulbs []int) (int, int) {
	// 每个color you两个位置
	pos := make([][]int, n)

	for i, c := range bulbs {
		c--
		pos[c] = append(pos[c], i)
	}

	cover := func(a, b []int) bool {
		if a[0] < b[0] && b[0] < a[1] {
			return true
		}
		if a[0] < b[1] && b[1] < a[1] {
			return true
		}
		return false
	}

	r := NewGraph(n, n*n)
	g := NewGraph(n, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if cover(pos[j], pos[i]) {
				g.AddEdge(i, j)
				r.AddEdge(j, i)
			}
		}
	}

	vis := make([]bool, n)

	var order []int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}

	belong := make([]int, n)
	for i := 0; i < n; i++ {
		belong[i] = -1
	}

	comp := make([][]int, n)

	var dfs2 func(u int, c int)

	dfs2 = func(u int, c int) {
		comp[c] = append(comp[c], u)
		belong[u] = c
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if belong[v] < 0 {
				dfs2(v, c)
			}
		}
	}
	var cid int
	for i := n - 1; i >= 0; i-- {
		u := order[i]
		if belong[u] < 0 {
			dfs2(u, cid)
			cid++
		}
	}

	deg := make([]int, cid)
	for u := 0; u < n; u++ {
		uc := belong[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			vc := belong[v]
			if uc != vc {
				deg[uc]++
			}
		}
	}

	var set int
	ways := 1

	for i := 0; i < cid; i++ {
		if deg[i] == 0 {
			set++
			ways = mul(ways, mul(2, len(comp[i])))
		}
	}

	return set, ways
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
