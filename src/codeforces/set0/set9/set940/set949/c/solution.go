package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, h := readThreeNums(reader)
	u := readNNums(reader, n)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, h, u, edges)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, i := range res {
		buf.WriteString(fmt.Sprintf("%d ", i))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(n int, h int, time []int, edges [][]int) []int {
	m := len(edges)

	g := NewGraph(n, m*2)
	rg := NewGraph(n, m*2)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if (time[u]+1)%h == time[v] {
			g.AddEdge(u, v)
			rg.AddEdge(v, u)
		}
		if (time[v]+1)%h == time[u] {
			g.AddEdge(v, u)
			rg.AddEdge(u, v)
		}
	}

	var order []int
	vis := make([]bool, n)
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

	comp := make([]int, n)
	for i := 0; i < n; i++ {
		comp[i] = -1
	}

	cnt := make([]int, n)

	var dfs2 func(u int, c int)

	dfs2 = func(u int, c int) {
		cnt[c]++
		comp[u] = c
		for i := rg.nodes[u]; i > 0; i = rg.next[i] {
			v := rg.to[i]
			if comp[v] < 0 {
				dfs2(v, c)
			}
		}
	}
	var cid int
	for i := n - 1; i >= 0; i-- {
		u := order[i]
		if comp[u] < 0 {
			dfs2(u, cid)
			cid++
		}
	}

	out := make([]int, cid)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if (time[u]+1)%h == time[v] && comp[u] != comp[v] {
			out[comp[u]]++
		}
		if (time[v]+1)%h == time[u] && comp[u] != comp[v] {
			out[comp[v]]++
		}
	}

	best := -1

	for i := 0; i < cid; i++ {
		if out[i] == 0 {
			if best < 0 || cnt[best] > cnt[i] {
				best = i
			}
		}
	}

	var arr []int

	for i := 0; i < n; i++ {
		if comp[i] == best {
			arr = append(arr, i+1)
		}
	}

	return arr
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
