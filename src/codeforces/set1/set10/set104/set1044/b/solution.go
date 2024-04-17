package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(x byte, y int) int {
		fmt.Printf("%c %d\n", x, y)
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		m := readNum(reader)
		A := readNNums(reader, m)
		m = readNum(reader)
		B := readNNums(reader, m)
		res := solve(n, edges, A, B, ask)
		fmt.Printf("C %d\n", res)
	}
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, edges [][]int, A []int, B []int, ask func(byte, int) int) int {
	g := NewGraph(n+1, 2*len(edges))

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	first := ask('B', B[0])

	for _, i := range A {
		if first == i {
			return first
		}
	}

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = -1
	}

	// 要找到离first最近的点
	que := make([]int, n)
	var front, tail int

	push := func(u int, d int) {
		que[front] = u
		front++
		dist[u] = d
	}

	pop := func() int {
		u := que[tail]
		tail++
		return u
	}
	push(first, 0)
	for tail < front {
		u := pop()
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				push(v, dist[u]+1)
			}
		}
	}

	second := A[0]
	for _, i := range A {
		if dist[i] < dist[second] {
			second = i
		}
	}
	ans := ask('A', second)

	for _, i := range B {
		if ans == i {
			return second
		}
	}
	return -1
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
