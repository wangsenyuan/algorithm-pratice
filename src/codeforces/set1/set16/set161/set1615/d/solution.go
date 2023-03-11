package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		X := make([][]int, m)
		for i := 0; i < m; i++ {
			X[i] = readNNums(reader, 3)
		}
		res := solve(n, E, X)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, e := range res {
				for i := 0; i < len(e); i++ {
					buf.WriteString(fmt.Sprintf("%d ", e[i]))
				}
				buf.WriteByte('\n')
			}
		}
	}
	fmt.Print(buf.String())
}

func solve(n int, E [][]int, X [][]int) [][]int {
	m := len(X)
	g := NewGraph(n, 2*n+2*m)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if w >= 0 {
			w = digit_count_xor(w)
			g.AddEdge(u, v, w)
			g.AddEdge(v, u, w)
		}
	}

	for _, x := range X {
		u, v, w := x[0]-1, x[1]-1, x[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	var dfs func(u int) bool
	dfs = func(u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if color[v] < 0 {
				color[v] = color[u] ^ g.val[i]
				if !dfs(v) {
					return false
				}
			} else if color[v] != color[u]^g.val[i] {
				return false
			}
		}
		return true
	}

	for u := 0; u < n; u++ {
		if color[u] < 0 {
			color[u] = 0
			if !dfs(u) {
				return nil
			}
		}
	}

	res := make([][]int, n-1)
	for i, cur := range E {
		if cur[2] >= 0 {
			res[i] = cur
		} else {
			u, v := cur[0], cur[1]
			w := color[u-1] ^ color[v-1]
			res[i] = []int{u, v, w}
		}
	}
	return res
}

func createGraphFromEdges(n int, E [][]int) *Graph {
	g1 := NewGraph(n, 2*n)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g1.AddEdge(u, v, w)
		g1.AddEdge(v, u, w)
	}
	return g1
}

func digit_count_xor(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res & 1
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
