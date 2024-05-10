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

	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

const mod = 1e9 + 7

const oo = 1 << 60

const X = 1e6

func solve(n int, roads [][]int) []int {
	g := NewGraph(X, X*2)
	n0 := n
	add := func(u, v, x int) {
		for x > 9 {
			g.AddEdge(n, v, x%10)
			v = n
			n++
			x /= 10
		}
		g.AddEdge(u, v, x)
	}

	for i, road := range roads {
		u, v := road[0]-1, road[1]-1
		add(u, v, i+1)
		add(v, u, i+1)
	}

	var que [][]int
	que = append(que, []int{0})
	vis := make([]bool, n)
	vis[0] = true

	dist := make([]int, n)

	for len(que) > 0 {
		tmp := que[0]
		que = que[1:]
		var buf [10][]int
		for _, u := range tmp {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				x := g.val[i]
				buf[x] = append(buf[x], u*n+v)
			}
		}
		for i, vs := range buf {
			var nxt []int
			for _, w := range vs {
				u, v := w/n, w%n
				if !vis[v] {
					dist[v] = (dist[u]*10 + i) % mod
					nxt = append(nxt, v)
					vis[v] = true
				}
			}
			if len(nxt) > 0 {
				que = append(que, nxt)
			}
		}
	}

	return dist[1:n0]
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
