package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	// tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		likes := make([][]int, m)
		for i := 0; i < m; i++ {
			likes[i] = readNNums(reader, 2)
		}
		res := solve(n, likes)

		if len(res) == 0 {
			buf.WriteString("No\n")
		} else {
			buf.WriteString(fmt.Sprintf("Yes\n%d %d\n", len(res[0]), len(res[1])))
			for _, cur := range res {
				for _, i := range cur {
					buf.WriteString(fmt.Sprintf("%d ", i))
				}
				buf.WriteByte('\n')
			}
		}

		if tc > 0 {
			readString(reader)
		}
	}

	fmt.Print(buf.String())
}

func solve(n int, likes [][]int) [][]int {
	// likes 是单向的
	m := len(likes)

	g := NewGraph(n, m)

	for _, like := range likes {
		u, v := like[0]-1, like[1]-1
		if u != v {
			g.AddEdge(u, v)
		}
	}

	comp := make([]int, n)
	var cid int

	low := make([]int, n)
	dfn := make([]int, n)
	var timer int
	stack := make([]int, n)
	var top int

	vis := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		stack[top] = u
		top++

		vis[u]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 2 {
				continue
			}
			if vis[v] == 1 {
				low[u] = min(low[u], dfn[v])
			} else {
				dfs(v)
				low[u] = min(low[u], low[v])
			}
		}

		if low[u] == dfn[u] {
			for top > 0 {
				v := stack[top-1]
				vis[v]++
				comp[v] = cid
				top--
				if u == v {
					break
				}
			}
			cid++
		}
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}

	if cid == 1 {
		return nil
	}

	res := make([][]int, 2)

	for i := 0; i < n; i++ {
		if comp[i] == 0 {
			res[0] = append(res[0], i+1)
		} else {
			res[1] = append(res[1], i+1)
		}
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
