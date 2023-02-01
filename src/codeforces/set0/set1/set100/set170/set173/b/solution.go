package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	maze := make([]string, n)
	for i := 0; i < n; i++ {
		maze[i] = readString(reader)[:m]
	}
	res := solve(maze)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(maze []string) int {
	n := len(maze)
	m := len(maze[0])

	g := NewGraph(n+m, 2*n*m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] == '#' {
				g.AddEdge(i, j+n)
				g.AddEdge(j+n, i)
			}
		}
	}

	que := make([]int, n+m)
	dist := make([]int, n+m)
	for i := 0; i < n+m; i++ {
		dist[i] = -1
	}
	dist[n-1] = 0

	var front, end int
	que[end] = n - 1
	end++

	for front < end {
		u := que[front]
		front++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[end] = v
				end++
			}
		}
	}

	return dist[0]
}
func solve1(maze []string) int {
	// dist[0][0]
	n := len(maze)
	m := len(maze[0])

	var tot int
	id := make([][]int, n)
	//var columns [][]int
	for i := 0; i < n; i++ {
		id[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if maze[i][j] == '#' {
				id[i][j] = tot
				tot++
			}
		}
	}

	vert := tot * 2
	// 每个单元格可以拆分成4个节点，分别代码上下左右,
	g := NewGraph(vert, 4*tot)

	//for i := 0; i < tot; i++ {
	//	x := i*2 + 0
	//	y := i*2 + 1
	//	// 转弯的cost
	//	g.AddEdge(x, y)
	//	g.AddEdge(y, x)
	//}

	prev := make([]int, m)
	for i := 0; i < m; i++ {
		prev[i] = -1
	}

	add := func(u []int, v []int, a int) {
		i := id[u[0]][u[1]]
		j := id[v[0]][v[1]]
		// 0 for horizontal, 1 for vertical
		g.AddEdge(i*2+a, j*2+a)
	}

	// # for column
	for i := 0; i < n; i++ {
		row := -1
		for j := 0; j < m; j++ {
			if maze[i][j] == '#' {
				if prev[j] != -1 {
					u := []int{prev[j], j}
					v := []int{i, j}
					add(u, v, 1)
					add(v, u, 1)
				}
				if row != -1 {
					u := []int{i, row}
					v := []int{i, j}
					add(u, v, 0)
					add(v, u, 0)
				}
				row = j
				prev[j] = i
			}
		}
	}

	que := make([]int, 2*vert)
	front, end := vert, vert
	dist := make([]int, vert)

	for i := 0; i < vert; i++ {
		dist[i] = -1
	}

	for j := m - 1; j >= 0; j-- {
		if maze[n-1][j] == '#' {
			x := id[(n - 1)][j] * 2
			dist[x] = 0
			que[end] = x
			end++
			break
		}
	}

	for front < end {
		u := que[front]
		front++

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]

			if dist[v] < 0 {
				dist[v] = dist[u]
				//y := v / 2
				// same node, but turn around
				front--
				que[front] = v
			}
		}
		// 0, 1, 2, 3
		x := u - 1
		if u%2 == 0 {
			x = u + 1
		}
		// x is another direction
		if dist[x] < 0 {
			dist[x] = dist[u] + 1
			que[end] = x
			end++
		}
	}

	res := -1

	for j := 0; j < m; j++ {
		if maze[0][j] == '#' {
			u := id[0][j]*2 + 0
			if dist[u] >= 0 {
				if res < 0 || dist[u] < res {
					res = dist[u]
				}
			}
		}
	}

	return res
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
