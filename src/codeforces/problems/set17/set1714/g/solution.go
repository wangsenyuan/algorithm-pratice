package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--

		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, E)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

const H = 20

func solve(n int, E [][]int) []int {
	g := NewGraph(n, 2*n)

	for i, e := range E {
		p, a, b := e[0], e[1], e[2]
		p--
		g.AddEdge(p, i+1, a, b)
	}

	ans := make([]int, n)

	var dfs func(u int, d int, a int, b int)

	stack := make([]int64, n)

	var sum_b int64
	var sum_a int64

	find := func(u int, d int) int {
		l, r := 0, d
		for l < r {
			mid := (l + r) / 2
			if stack[mid] > sum_a {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r - 1
	}

	dfs = func(u int, d int, a, b int) {
		sum_a += int64(a)
		sum_b += int64(b)
		stack[d] = sum_b
		ans[u] = find(u, d+1)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v, d+1, g.a[i], g.b[i])
		}
		sum_a -= int64(a)
		sum_b -= int64(b)
	}

	dfs(0, 0, 0, 0)

	return ans[1:]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	a     []int
	b     []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	a := make([]int, e+1)
	b := make([]int, e+1)
	return &Graph{nodes, next, to, a, b, 0}
}

func (g *Graph) AddEdge(u, v, a, b int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.a[g.cur] = a
	g.b[g.cur] = b
}
