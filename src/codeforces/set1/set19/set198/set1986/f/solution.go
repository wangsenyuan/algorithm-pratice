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
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, edges)
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

func solve(n int, edges [][]int) int {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	// sz[u] = 从u能访问到的节点的数量
	belong := make([]int, n)
	for i := 0; i < n; i++ {
		belong[i] = -1
	}
	var comp []int

	var dfs func(u int, c int) int

	dfs = func(u int, c int) int {
		belong[u] = c
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if belong[v] < 0 {
				sz += dfs(v, c)
			}
		}
		return sz
	}

	var tot int

	for i := 0; i < n; i++ {
		if belong[i] < 0 {
			sz := dfs(i, len(comp))
			comp = append(comp, sz)
			tot += sz * (sz - 1) / 2
		}
	}
	vis := make([]bool, n)
	dis := make([]int, n)
	low := make([]int, n)

	var time int

	ans := tot

	var dfs2 func(p int, u int) int
	dfs2 = func(p int, u int) int {
		dis[u] = time
		low[u] = time
		time++

		vis[u] = true
		res := 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				// 到parent来的不算
				continue
			}
			if !vis[v] {
				tmp := dfs2(u, v)
				low[u] = min(low[u], low[v])
				if low[v] > dis[u] {
					// a bridge
					sz := comp[belong[u]] - tmp
					ans = min(ans, tot-sz*tmp)
				}
				res += tmp
			} else {
				// a back edge
				low[u] = min(low[u], dis[v])
			}
		}

		return res
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs2(-1, i)
		}
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
