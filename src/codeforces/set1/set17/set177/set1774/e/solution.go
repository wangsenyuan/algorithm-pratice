package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, d := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	s, _ := reader.ReadBytes('\n')
	var cnt int
	pos := readInt(s, 0, &cnt)
	a := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		pos = readInt(s, pos+1, &a[i])
	}
	s, _ = reader.ReadBytes('\n')
	pos = readInt(s, 0, &cnt)
	b := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		pos = readInt(s, pos+1, &b[i])
	}
	res := solve(n, E, a, b, d)
	fmt.Println(res)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, E [][]int, a []int, b []int, d int) int {
	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	anc := make([]int, n)
	path := make([]int, n)

	var dfs func(p int, u int, dep int)

	dfs = func(p int, u int, dep int) {
		path[dep] = u
		anc[u] = -1
		if dep >= d {
			anc[u] = path[dep-d]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, dep+1)
			}
		}
	}

	dfs(0, 0, 0)

	f := make([][]bool, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([]bool, n)
	}

	for _, x := range a {
		x--
		f[0][x] = true
		if anc[x] >= 0 {
			f[1][anc[x]] = true
		}
	}

	for _, x := range b {
		x--
		f[1][x] = true
		if anc[x] >= 0 {
			f[0][anc[x]] = true
		}
	}

	var dfs2 func(p int, u int, tp int) bool

	dfs2 = func(p int, u int, tp int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if dfs2(u, v, tp) {
					f[tp][u] = true
				}
			}
		}
		return f[tp][u]
	}

	for i := 0; i < 2; i++ {
		dfs2(-1, 0, i)
	}
	var ans int
	for i := 0; i < 2; i++ {
		for j := 1; j < n; j++ {
			if f[i][j] {
				ans += 2
			}
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
