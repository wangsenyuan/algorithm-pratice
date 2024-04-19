package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	edges := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res, ops := solve(n, edges)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", res))

	for _, op := range ops {
		for i := 0; i < len(op); i++ {
			buf.WriteString(fmt.Sprintf("%d ", op[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, edges [][]int) (int, [][]int) {
	g := NewGraph(n, n*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)

	// get the distance from x, and return the farthest vertex
	bfs := func(x int) (int, []int, []int) {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		track := make([]int, n)
		var front, tail int
		que[front] = x
		front++
		dist[x] = 0
		for tail < front {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					track[v] = u
					dist[v] = dist[u] + 1
					que[front] = v
					front++
				}
			}
		}
		res := x
		for i := 0; i < n; i++ {
			if dist[i] > dist[res] {
				res = i
			}
		}
		return res, dist, track
	}

	first, _, _ := bfs(0)
	second, dist1, track := bfs(first)
	_, dist2, _ := bfs(second)

	dia := make([]bool, n)
	for i := second; i != first; i = track[i] {
		dia[i] = true
	}
	dia[first] = true

	var res int
	var ops [][]int
	var other []int
	for i := 0; i < n; i++ {
		if !dia[i] {
			other = append(other, i)
		}
	}

	sort.Slice(other, func(i, j int) bool {
		return dist1[other[i]] > dist1[other[j]]
	})

	for _, i := range other {
		if dist1[i] >= dist2[i] {
			ops = append(ops, []int{first + 1, i + 1, i + 1})
			res += dist1[i]
		} else {
			ops = append(ops, []int{second + 1, i + 1, i + 1})
			res += dist2[i]
		}
	}

	for i := second; i != first; i = track[i] {
		ops = append(ops, []int{first + 1, i + 1, i + 1})
		res += dist1[i]
	}

	return res, ops
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
