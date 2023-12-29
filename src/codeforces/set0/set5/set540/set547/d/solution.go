package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	pts := make([][]int, n)
	for i := 0; i < n; i++ {
		pts[i] = readNNums(reader, 2)
	}

	res := solve(pts)

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

func solve(pts [][]int) string {
	n := len(pts)

	// 每个点最多4条边
	g := NewGraph(n, 4*n)

	points := make([]Point, n)

	for i := 0; i < n; i++ {
		points[i] = Point{i, pts[i][0], pts[i][1]}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x || points[i].x == points[j].x && points[i].y < points[j].y
	})

	for i := 0; i < n; {
		j := i
		for i < n && points[i].x == points[j].x {
			if (i-j)%2 == 1 {
				u := points[i-1].id
				v := points[i].id
				g.AddEdge(u, v)
				g.AddEdge(v, u)
			}
			i++
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y || points[i].y == points[j].y && points[i].x < points[j].x
	})

	for i := 0; i < n; {
		j := i
		for i < n && points[i].y == points[j].y {
			if (i-j)%2 == 1 {
				u := points[i-1].id
				v := points[i].id
				g.AddEdge(u, v)
				g.AddEdge(v, u)
			}
			i++
		}
	}

	color := make([]int, n)

	var dfs func(u int, c int)
	dfs = func(u int, c int) {
		color[u] = c
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if color[v] == 0 {
				dfs(v, -c)
			}
		}
	}

	for i := 0; i < n; i++ {
		if color[i] == 0 {
			dfs(i, 1)
		}
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		if color[i] == 1 {
			res[i] = 'r'
		} else {
			res[i] = 'b'
		}
	}
	return string(res)
}

type Point struct {
	id int
	x  int
	y  int
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
