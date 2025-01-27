package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (f []int, res [][]int) {
	n := readNum(reader)
	f = readNNums(reader, n)
	res = solve(f)
	return
}

func solve(f []int) [][]int {
	n := len(f)

	adj := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		p := f[i] - 1
		adj[i] = append(adj[i], p)
		deg[p]++
	}

	var end []int
	var begin []int

	vis := make([]bool, n)
	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for _, v := range adj[u] {
			if vis[v] {
				end = append(end, v)
				return
			}
			dfs(v)
		}
	}

	var leaves int
	for u := 0; u < n; u++ {
		if deg[u] == 0 {
			leaves++
			begin = append(begin, u)
			dfs(u)
		}
	}

	for u := 0; u < n; u++ {
		if !vis[u] {
			begin = append(begin, u)
			dfs(u)
		}
	}
	if len(begin) == 1 && leaves == 0 {
		return nil
	}
	var res [][]int
	m := len(begin)
	for i := 0; i < m; i++ {
		res = append(res, []int{end[i] + 1, begin[(i+1)%m] + 1})
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
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
