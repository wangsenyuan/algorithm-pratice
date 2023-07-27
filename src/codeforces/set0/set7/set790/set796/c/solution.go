package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E, A)
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

func solve(n int, E [][]int, A []int) int {
	// max(A) + 2 or max(A) or max(A) + 1
	first := max(A...)

	var critical []int
	for i := 0; i < n; i++ {
		if A[i] == first {
			critical = append(critical, i)
		}
	}

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	que := make([]int, n)
	var front, end int
	que[end] = critical[0]
	end++
	dist[que[front]] = 0

	par := make([]int, n)
	for front < end {
		u := que[front]
		front++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				par[v] = u
				dist[v] = dist[u] + 1
				que[end] = v
				end++
			}
		}
	}
	center := critical[0]
	for i := 0; i < n; i++ {
		if A[i] == first {
			if dist[i] > 2 {
				return first + 2
			}
			if dist[i] == 2 {
				center = par[i]
			}
		}
	}

	var dfs func(p int, u int, expect int) bool

	dfs = func(p int, u int, expect int) bool {
		var add int
		if u == center {
			add = 0
		} else if p == center {
			add = 1
		} else {
			add = 2
		}
		if A[u]+add > expect {
			return false
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !dfs(u, v, expect) {
				return false
			}
		}
		return true
	}

	if dfs(-1, center, first) {
		return first
	}
	if dfs(-1, center, first+1) {
		return first + 1
	}

	return first + 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
