package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		C := make([][]int, m)
		for i := 0; i < m; i++ {
			C[i] = readNNums(reader, 3)
		}
		res := solve(n, m, C)
		if res > 0 {
			buf.WriteString(fmt.Sprintf("%d\n", res))
		} else {
			buf.WriteString("Impossible\n")
		}
	}
	fmt.Print(buf.String())
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	e += 10
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

func (g *Graph) Reset() {
	g.cur = 0
	for i := 0; i < len(g.nodes); i++ {
		g.nodes[i] = 0
	}
}

const INF = 1000000000

func solve(n, m int, C [][]int) int {
	g := NewGraph(n, m)

	disc := make([]int, n)
	low := make([]int, n)
	id := make([]int, n)
	vis := make([]bool, n)
	stack := make([]int, n)
	var p int
	var dfs func(u int, time *int, ii *int)

	dfs = func(u int, time *int, ii *int) {
		disc[u] = *time
		low[u] = disc[u]
		*time++
		stack[p] = u
		p++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if disc[v] < 0 {
				dfs(v, time, ii)
				low[u] = min(low[u], low[v])
			} else {
				low[u] = min(low[u], disc[v])
			}
		}

		if disc[u] == low[u] {
			// a head
			for p > 0 {
				v := stack[p-1]
				id[v] = *ii
				p--
				if v == u {
					break
				}
			}
			*ii++
		}
	}

	g2 := NewGraph(n, m)
	d := make([]int, n)
	var dfs2 func(u int)

	dfs2 = func(u int) {
		if vis[u] {
			return
		}
		vis[u] = true
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			dfs2(v)
			d[u] = max(d[u], d[v])
		}
		d[u]++
	}

	check := func(x int) bool {
		g.Reset()
		for i := 0; i < m; i++ {
			c := C[i]
			if c[2] <= x {
				g.AddEdge(c[0]-1, c[1]-1)
			}
		}

		for i := 0; i < n; i++ {
			disc[i] = -1
			d[i] = 0
			vis[i] = false
		}

		p = 0

		var scc int
		time := new(int)
		for i := 0; i < n; i++ {
			if disc[i] < 0 {
				dfs(i, time, &scc)
			}
		}

		g2.Reset()

		for u := 0; u < n; u++ {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if id[u] != id[v] {
					g2.AddEdge(id[u], id[v])
				}
			}
		}

		for i := 0; i < scc; i++ {
			if !vis[i] {
				dfs2(i)
			}
			if d[i] == scc {
				return true
			}
		}
		return false
	}

	if !check(INF) {
		return -1
	}
	left, right := 1, INF
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
