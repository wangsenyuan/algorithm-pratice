package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, s := readThreeNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}
	x, a, y, b := solve(n, s, edges)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n%s\n%d\n%s\n", x, a, y, b))
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, s int, edges [][]int) (int, string, int, string) {
	g := NewGraph(n, len(edges)*2)

	for i, e := range edges {
		u, v := e[1]-1, e[2]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	x, a := solveMax(s-1, n, g, edges)
	y, b := solveMin(s-1, n, g, edges)

	return x, a, y, b
}

func solveMin(s int, n int, g *Graph, edges [][]int) (int, string) {
	que := make([]int, n)
	vis := make([]bool, n)

	front, tail := 0, 0
	que[front] = s
	front++
	vis[s] = true
	ans := make([]int, len(edges))

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			e := edges[j]
			if vis[v] {
				// already visited
				continue
			}
			if e[0] == 1 {
				// a directed edge
				if u == e[1]-1 {
					vis[v] = true
					que[front] = v
					front++
				}
				// else do nothing
			} else {
				// a bi-directed edge
				if u == e[1]-1 {
					// 正向
					ans[j] = -1
				} else {
					// 反向
					ans[j] = 1
				}
			}
		}
	}

	var buf bytes.Buffer

	for i, e := range edges {
		if e[0] == 2 {
			if ans[i] == 1 {
				buf.WriteString("+")
			} else {
				buf.WriteString("-")
			}
		}
	}
	return front, buf.String()
}

func solveMax(s int, n int, g *Graph, edges [][]int) (int, string) {
	que := make([]int, n)
	vis := make([]bool, n)

	front, tail := 0, 0
	que[front] = s
	front++
	vis[s] = true
	ans := make([]int, len(edges))

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			e := edges[j]
			if vis[v] {
				// already visited
				continue
			}
			if e[0] == 1 {
				// a directed edge
				if u == e[1]-1 {
					vis[v] = true
					que[front] = v
					front++
				}
				// else do nothing
			} else {
				// a bi-directed edge
				if u == e[1]-1 {
					// 正向
					ans[j] = 1
				} else {
					// 反向
					ans[j] = -1
				}
				vis[v] = true
				que[front] = v
				front++
			}
		}
	}

	var buf bytes.Buffer

	for i, e := range edges {
		if e[0] == 2 {
			if ans[i] == 1 {
				buf.WriteString("+")
			} else {
				buf.WriteString("-")
			}
		}
	}
	return front, buf.String()
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
