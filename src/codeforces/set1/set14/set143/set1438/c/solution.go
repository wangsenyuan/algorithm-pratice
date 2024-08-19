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
		n, m := readTwoNums(reader)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = readNNums(reader, m)
		}
		res := solve1(a)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
			}
			buf.WriteByte('\n')
		}
	}

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

func solve1(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i+j)%2 == 1 {
				// 把这里变成奇数
				if a[i][j]%2 == 0 {
					a[i][j]++
				}
			} else {
				if a[i][j]%2 == 1 {
					a[i][j]++
				}
			}
		}
	}
	return a
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	// i * m + j => v
	// i * m + j + m * n => ~v
	V := n * m
	g := NewGraph(2*V, 8*V)
	gr := NewGraph(2*V, 8*V)

	add := func(u int, v int) {
		g.AddEdge(u, v)
		gr.AddEdge(v, u)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				u, v := i+dd[k], j+dd[k+1]
				if u < 0 || u == n || v < 0 || v == m {
					continue
				}
				if a[i][j]+1 == a[u][v] {
					add(i*m+j, u*m+v)
					add(u*m+v+V, i*m+j+V)
				}
				if a[i][j] == a[u][v] {
					add(i*m+j+V, u*m+v)
					add(u*m+v+V, i*m+j)
					add(u*m+v, i*m+j+V)
					add(i*m+j, u*m+v+V)
				}
			}
		}
	}

	vis := make([]bool, 2*V)

	order := make([]int, 2*V)
	var top int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order[top] = u
		top++
	}

	for i := 0; i < 2*V; i++ {
		if !vis[i] {
			dfs(i)
		}
	}

	comp := make([]int, 2*V)

	var dfs2 func(u int, c int)
	dfs2 = func(u int, c int) {
		comp[u] = c
		for i := gr.nodes[u]; i > 0; i = gr.next[i] {
			v := gr.to[i]
			if comp[v] == 0 {
				dfs2(v, c)
			}
		}
	}
	var c int
	for top > 0 {
		u := order[top-1]
		top--
		if comp[u] == 0 {
			c++
			dfs2(u, c)
		}
	}

	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		copy(res[i], a[i])

		for j := 0; j < m; j++ {
			if comp[i*m+j] == comp[i*m+j+V] {
				return nil
			}
			if comp[i*m+j] > comp[i*m+j+V] {
				res[i][j]++
			}
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
