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
		n, m, k := readThreeNums(reader)
		A := make([][]int, m)
		for i := 0; i < m; i++ {
			A[i] = readNNums(reader, 2)
		}
		B := make([][]int, k)
		for i := 0; i < k; i++ {
			B[i] = readNNums(reader, 3)
		}
		res := solve(n, A, B)
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
		if s[i] == '\n' {
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

func solve(n int, A [][]int, B [][]int) bool {
	// abs(A[u] - A[v]) <= d
	// 假设A[u] >= A[v]
	// A[u] - A[v] <= d
	//  A[u] >= A[v]
	//  A[u] <= A[v] + d
	g := NewGraph(n+1, (len(A)+len(B))*2)

	for _, a := range A {
		x, v := a[0], a[1]
		g.AddEdge(0, x, -v)
		g.AddEdge(x, 0, v)
	}

	for _, b := range B {
		u, v, d := b[0], b[1], b[2]
		g.AddEdge(u, v, d)
		g.AddEdge(v, u, d)
	}

	dist := make([]int64, n+1)

	for it := 0; it < n+1; it++ {
		change := false
		for u := 0; u <= n; u++ {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if dist[v] > dist[u]+int64(w) {
					dist[v] = dist[u] + int64(w)
					change = true
				}
			}
		}
		if !change {
			return true
		}
	}
	return false
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
	value := make([]int, e+1)
	return &Graph{nodes, next, to, value, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
