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
		n, _ := readTwoNums(reader)
		a := readNNums(reader, n)
		h := readNNums(reader, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(a, h, edges)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(a []int, h []int, edges [][]int) bool {
	n := len(a)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	type result struct {
		ok   bool
		good int
		bad  int
	}

	fail := func() result {
		return result{false, -1, -1}
	}

	sucess := func(x int, y int) result {
		return result{true, x, y}
	}

	sz := make([]int, n)
	copy(sz, a)

	var dfs func(p int, u int) result

	dfs = func(p int, u int) result {
		// x for good mood, y for bad modd
		var x, y int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v)
				if !tmp.ok {
					return fail()
				}
				sz[u] += sz[v]
				x += tmp.good
				y += tmp.bad
			}
		}
		// 那些坏心情的，他们在这个city，有可能是好心情的
		// 但是好心情的，仍然是好心情的
		// 假设这里分别有 w, v 个人
		// w + v = sz[u]
		// w - v = h[u]
		if (sz[u]+h[u]) < 0 || (sz[u]+h[u])%2 != 0 {
			return fail()
		}
		w := (sz[u] + h[u]) / 2
		v := sz[u] - w

		// 好心情的会变坏，所以w会变小, v 会变大
		if v < 0 || w < x {
			return fail()
		}

		return sucess(w, v)
	}

	res := dfs(-1, 0)

	return res.ok
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
