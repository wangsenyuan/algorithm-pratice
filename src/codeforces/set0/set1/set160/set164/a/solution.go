package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	V := readNNums(reader, n)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, V, E)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

func solve(n int, V []int, E [][]int) []int {
	g1 := NewGraph(n, len(E))
	g2 := NewGraph(n, len(E))
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g1.AddEdge(u, v)
		g2.AddEdge(v, u)
	}

	v1 := make([]bool, n)

	var dfs1 func(u int)

	dfs1 = func(u int) {
		v1[u] = true
		for i := g1.node[u]; i > 0; i = g1.next[i] {
			v := g1.to[i]
			if !v1[v] {
				dfs1(v)
			}
		}
	}

	v2 := make([]bool, n)

	var dfs2 func(u int)
	dfs2 = func(u int) {
		v2[u] = true
		if V[u] != 1 {
			for i := g2.node[u]; i > 0; i = g2.next[i] {
				v := g2.to[i]
				if !v2[v] {
					dfs2(v)
				}
			}
		}

	}

	for i := 0; i < n; i++ {
		if V[i] == 1 && !v1[i] {
			dfs1(i)
		}
	}

	for i := 0; i < n; i++ {
		if V[i] == 2 && !v2[i] {
			dfs2(i)
		}
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if v1[i] && v2[i] {
			ans[i] = 1
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
