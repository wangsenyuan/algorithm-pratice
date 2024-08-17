package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 3)
	}
	res := solve(n, edges)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, edges [][]int) int64 {
	g := NewGraph(n, 2*n)
	for i, cur := range edges {
		u, v := cur[0], cur[1]
		g.AddEdge(u-1, v-1, i)
		g.AddEdge(v-1, u-1, i)
	}

	sz := make([]int, n)
	se := make([]int, n-1)

	fa := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sz[u] = 1
		fa[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				se[j] = min2(sz[v], n-sz[v])
			}
		}
	}

	dfs(-1, 0)

	var sum int64

	var res int64

	for i := 0; i < n-1; i++ {
		sum += 2 * int64(se[i]) * int64(edges[i][2])
	}

	for i := 0; i < n-1; i++ {
		if se[i]*2 == n {
			res = max(res, sum-int64(edges[i][2]))
		}
	}

	for u := 0; u < n; u++ {
		ok := true
		var cnt int
		me := 1 << 30
		for i := g.nodes[u]; i > 0 && ok; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if v != fa[u] {
				cnt += sz[v]
				ok = sz[v]*2 < n
			}
			if edges[j][2] < me {
				me = edges[j][2]
			}
		}

		if ok && (n-cnt-1)*2 < n {
			// found
			res = max(res, sum-int64(me))
		}
	}

	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
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
