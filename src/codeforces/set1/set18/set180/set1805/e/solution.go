package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	A := readNNums(reader, n)
	res := solve(n, A, E)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(n int, A []int, E [][]int) []int {
	total := make(map[int]int)

	var mad int

	for _, x := range A {
		total[x]++
		if total[x] >= 2 {
			mad = max(mad, x)
		}
	}

	ans := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		ans[i] = -1
	}
	if mad == 0 {
		// no one exceeds 2
		for i := 0; i < n-1; i++ {
			ans[i] = 0
		}
		return ans
	}

	if total[mad] > 2 {
		// all mad
		for i := 0; i < len(ans); i++ {
			ans[i] = mad
		}
		return ans
	}
	g := NewGraph(n, 2*n)

	for i, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}
	first, second := -1, -1
	next := make([]int, n)

	for i := 0; i < n; i++ {
		if A[i] == mad {
			if first < 0 {
				first = i
			} else {
				second = i
			}
		}
		next[i] = -1
	}

	var dfs func(p int, u int) bool

	dfs = func(p int, u int) bool {
		if u == second {
			return true
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && dfs(u, v) {
				next[u] = i
				return true
			}
		}
		return false
	}

	dfs(-1, first)

	// go from first to second
	// 先计算出不包括next[i]中的最大值
	var secondMad int
	var thirdMad int

	var dfs2 func(p int, u int)

	var dfs3 func(p int, u int)

	var dfs4 func(p int, u int)

	freq3 := make(map[int]int)

	dfs3 = func(p int, u int) {
		freq3[A[u]]++
		if freq3[A[u]] >= 2 && A[u] >= secondMad {
			secondMad = A[u]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs3(u, v)
			}
		}
	}

	freq2 := make(map[int]int)

	dfs4 = func(p int, u int) {
		freq2[A[u]]++
		if freq2[A[u]] >= 2 && A[u] >= thirdMad {
			thirdMad = A[u]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs4(u, v)
			}
		}
	}

	dfs2 = func(p int, u int) {
		if u == second {
			dfs4(p, u)
			return
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && i != next[u] {
				dfs3(u, v)
			}
		}

		freq3[A[u]]++
		if freq3[A[u]] >= 2 && A[u] >= secondMad {
			secondMad = A[u]
		}

		v := g.to[next[u]]
		ans[g.val[next[u]]] = secondMad
		dfs2(u, v)
		ans[g.val[next[u]]] = max(ans[g.val[next[u]]], thirdMad)

		freq2[A[u]]++
		if freq2[A[u]] >= 2 && A[u] >= thirdMad {
			thirdMad = A[u]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			w := g.to[i]
			if p != w && i != next[u] {
				dfs4(u, w)
			}
		}
	}

	dfs2(-1, first)

	for i := 0; i < n-1; i++ {
		if ans[i] < 0 {
			ans[i] = mad
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
