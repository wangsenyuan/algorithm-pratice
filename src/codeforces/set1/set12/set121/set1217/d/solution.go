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
	n, k := readTwoNums(reader)
	edges := make([][]int, k)
	for i := 0; i < k; i++ {
		edges[i] = readNNums(reader, 2)
	}

	cnt, res := solve(n, edges)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", cnt))
	for i := 0; i < k; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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

func solve(n int, edges [][]int) (int, []int) {
	g := NewGraph(n, len(edges))

	for i, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1, i)
	}
	vis := make([]int, n)

	var dfs func(u int)
	color := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		color[i] = 1
	}

	dfs = func(u int) {
		vis[u]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]

			if vis[v] == 0 {
				dfs(v)
			} else if vis[v] == 1 {
				color[g.val[i]] = 2
			}
		}

		vis[u]++
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}

	return slices.Max(color), color
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
	e++
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
