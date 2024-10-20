package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = readNNums(reader, 2)
	}

	ok, res := solve(n, roads)

	if !ok {
		fmt.Println("Impossible")
	} else {
		fmt.Println(res)
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

func solve(n int, roads [][]int) (bool, string) {
	m := len(roads)

	g := NewGraph(m, m*m*2)

	for i := 0; i < m; i++ {
		u, v := roads[i][0], roads[i][1]
		u, v = min(u, v), max(u, v)
		for j := i + 1; j < m; j++ {
			x, y := roads[j][0], roads[j][1]
			x, y = min(x, y), max(x, y)
			if u < x && x < v && v < y || x < u && u < y && y < v {
				// overlap
				g.AddEdge(i, j)
				g.AddEdge(j, i)
			}
		}
	}

	color := make([]int, m)
	vis := make([]bool, m)

	var dfs func(u int, c int) bool

	dfs = func(u int, c int) bool {
		if vis[u] {
			return color[u] == c
		}
		vis[u] = true
		color[u] = c
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, c^1) {
				return false
			}
		}
		return true
	}

	for i := 0; i < m; i++ {
		if !vis[i] {
			if !dfs(i, 0) {
				return false, ""
			}
		}
	}

	ans := make([]byte, m)
	for i := 0; i < m; i++ {
		if color[i] == 1 {
			ans[i] = 'i'
		} else {
			ans[i] = 'o'
		}
	}

	return true, string(ans)
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
