package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	rounds := make([]int, m)
	for i := 0; i < m; i++ {
		rounds[i] = readNum(reader)
	}
	ans := solve(n, edges, rounds)

	var buf bytes.Buffer

	for _, x := range ans {
		buf.WriteString(x)
		buf.WriteByte('\n')
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

const Ron = "Ron"
const Hermione = "Hermione"

func solve(n int, edges [][]int, rounds []int) []string {
	deg := make([]int, n)
	g := NewGraph(n, n*2)
	for _, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		deg[u]++
		deg[v]++
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	first := 0
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			first = i
			break
		}
	}

	pos := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				pos[v] = pos[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(-1, first)

	ans := make([]string, len(rounds))

	for i, u := range rounds {
		u--
		x := pos[u]
		y := n - 1 - x
		if x&1 == 1 || y&1 == 1 {
			ans[i] = Ron
		} else {
			ans[i] = Hermione
		}
	}

	return ans
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
