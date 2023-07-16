package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
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

const inf = 1 << 30

func solve(n int, roads [][]int) int {

	g := NewGraph(n*2, len(roads))
	var mw int
	for _, r := range roads {
		u, v, w := r[0], r[1], r[2]
		u--
		v--
		g.AddEdge(u, v+n, w)
		if mw < w {
			mw = w
		}
	}

	pu := make([]int, n+1)
	pv := make([]int, n+1)
	dist := make([]int, n+1)
	que := make([]int, n+1)
	bfs := func(d int) bool {
		var front, end int
		for i := 0; i < n; i++ {
			if pu[i] < 0 {
				dist[i] = 0
				que[end] = i
				end++
			} else {
				dist[i] = inf
			}
		}
		dist[n] = inf
		for front < end {
			u := que[front]
			front++
			if u == n {
				continue
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i] - n
				w := g.val[i]
				if w <= d && dist[pv[v]] == inf {
					dist[pv[v]] = dist[u] + 1
					que[end] = pv[v]
					end++
				}
			}
		}
		return dist[n] < inf
	}

	var dfs func(u int, d int) bool

	dfs = func(u int, d int) bool {
		if u == n {
			return true
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i] - n
			w := g.val[i]
			if w > d {
				continue
			}
			if dist[pv[v]] == dist[u]+1 && dfs(pv[v], d) {
				pv[v] = u
				pu[u] = v
				return true
			}
		}
		dist[u] = inf
		return false
	}

	check := func(d int) bool {
		for i := 0; i < n; i++ {
			pu[i] = -1
			pv[i] = n
			dist[i] = inf
		}
		var cnt int

		for bfs(d) && cnt < n {
			for i := 0; i < n; i++ {
				if pu[i] < 0 && dfs(i, d) {
					cnt++
				}
			}
		}
		return cnt == n
	}

	if !check(mw) {
		return -1
	}

	l, r := 0, mw
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

func (g *Graph) Reset() {
	for i := 0; i < len(g.nodes); i++ {
		g.nodes[i] = 0
	}
	g.cur = 0
}
