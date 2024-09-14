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
		edges[i] = readNNums(reader, 2)
	}

	cnt, res := solve(n, edges)

	fmt.Println(cnt)
	fmt.Println(res[0], res[1], res[2])
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

func solve(n int, edges [][]int) (int, []int) {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([]int, n)
	que := make([]int, n)
	fa := make([]int, n)
	bfs := func(s []int) {
		for i := 0; i < n; i++ {
			dist[i] = -1
			fa[i] = -1
		}
		var head, tail int
		for _, x := range s {
			dist[x] = 0
			que[head] = x
			head++
		}

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					fa[v] = u
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
	}

	bfs([]int{0})
	var first int
	for i := 0; i < n; i++ {
		if dist[first] < dist[i] {
			first = i
		}
	}

	bfs([]int{first})
	var second int
	for i := 0; i < n; i++ {
		if dist[second] < dist[i] {
			second = i
		}
	}

	ans := dist[second]

	// (first -> .. -> second) 是直径
	var dia []int
	for i := second; i != -1; i = fa[i] {
		dia = append(dia, i)
	}

	bfs(dia)

	third := -1
	for i := 0; i < n; i++ {
		if i == first || i == second {
			continue
		}
		if third < 0 || dist[i] > dist[third] {
			third = i
		}
	}

	ans += dist[third]

	return ans, []int{first + 1, second + 1, third + 1}
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
