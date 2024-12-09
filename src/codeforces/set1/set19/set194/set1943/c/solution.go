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
		edges := make([][]int, n-1)
		for i := range n - 1 {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, edges)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(n int, edges [][]int) [][]int {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)

	bfs := func(x int) ([]int, int) {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		var head, tail int
		que[head] = x
		head++
		dist[x] = 0
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}

		res := x
		for i := 0; i < n; i++ {
			if dist[i] > dist[res] {
				res = i
			}
		}

		return dist, res
	}

	_, a := bfs(0)
	d2, b := bfs(a)
	d3, c := bfs(b)

	diam := d3[c] + 1

	var centers []int
	for i := 0; i < n; i++ {
		if d2[i]+d3[i] == diam-1 && (d2[i] == diam/2 || d3[i] == diam/2) {
			centers = append(centers, i)
		}
	}
	var res [][]int

	if len(centers) == 1 {
		for i := 0; i <= diam/2; i++ {
			res = append(res, []int{centers[0] + 1, i})
		}
	} else {
		x, y := centers[0]+1, centers[1]+1
		need := diam/2 - 1
		for need >= 0 {
			res = append(res, []int{x, need})
			res = append(res, []int{y, need})
			need -= 2
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
