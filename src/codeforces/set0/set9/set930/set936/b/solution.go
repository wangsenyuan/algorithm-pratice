package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res, path := process(reader)
	if res != "Win" {
		fmt.Println(res)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(res)
	buf.WriteByte('\n')
	for _, x := range path {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func process(reader *bufio.Reader) (n int, edges [][]int, s int, res string, path []int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, n)
	for i := range n {
		var k int
		bs, _ := reader.ReadBytes('\n')
		pos := readInt(bs, 0, &k) + 1
		edges[i] = make([]int, k)
		for j := range k {
			pos = readInt(bs, pos, &edges[i][j]) + 1
		}
	}
	s = readNum(reader)
	res, path = solve(n, m, edges, s)
	return
}

func solve(n int, m int, edges [][]int, s int) (string, []int) {
	s--
	g := NewGraph(n*2, 2*m)
	g2 := NewGraph(n, m)
	deg := make([]int, n)

	add := func(u, v int) {
		deg[u]++
		g.AddEdge(2*u, 2*v+1)
		g.AddEdge(2*u+1, 2*v)
		g2.AddEdge(u, v)
	}

	for i, cur := range edges {
		for _, j := range cur {
			add(i, j-1)
		}
	}

	que := make([]int, n*2)
	vis := make([]bool, n*2)

	var head, tail int
	push := func(u int) {
		vis[u] = true
		que[head] = u
		head++
	}

	poll := func() (u int) {
		u = que[tail]
		tail++
		return
	}

	push(2*s + 1)

	trace := make([]int, 2*n)
	for i := range 2 * n {
		trace[i] = -1
	}

	for tail < head {
		u := poll()
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				trace[v] = u
				push(v)
			}
		}
	}

	for i := 0; i < n; i++ {
		if vis[2*i] && deg[i] == 0 {
			var path []int
			x := 2 * i
			for x >= 0 {
				path = append(path, x/2+1)
				x = trace[x]
			}
			reverse(path)
			return "Win", path
		}
	}

	marked := make([]int, n)

	var dfs func(u int) bool

	dfs = func(u int) bool {
		if marked[u] == 1 {
			return true
		}
		if marked[u] == 2 {
			return false
		}
		marked[u] = 1
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			if dfs(v) {
				return true
			}
		}
		marked[u]++
		return false
	}

	if dfs(s) {
		return "Draw", nil
	}

	return "Lose", nil
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
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
