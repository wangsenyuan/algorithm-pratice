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
		p := readNum(reader)
		E := make([][]int, (1<<p)-1)
		for i := 0; i < len(E); i++ {
			E[i] = readNNums(reader, 2)
		}
		r, v, x := solve(p, E)

		buf.WriteString(fmt.Sprintf("%d\n", r))
		for i := 0; i < len(v); i++ {
			buf.WriteString(fmt.Sprintf("%d ", v[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < len(x); i++ {
			buf.WriteString(fmt.Sprintf("%d ", x[i]))
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

func solve(m int, E [][]int) (int, []int, []int) {
	// 1 ... 2 * n - 1
	// f(root) 表示为给定root的结果，假设 v(x) = 节点x的赋值
	// f(root) >= v(root)
	// 如果 v(root) 的 最高位是 h(root) 果然就是 h(root)
	n := len(E) + 1
	// 1 << p = n
	ans := make([]int, n)
	edge := make([]int, n-1)

	g := NewGraph(n, 2*n)

	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	var x int

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if p != v {
				x++
				if ans[u]&(1<<m) > 0 {
					ans[v] = x
					edge[w] = x + (1 << m)
				} else {
					ans[v] = x + (1 << m)
					edge[w] = x
				}
				dfs(u, v)
			}
		}
	}

	ans[0] = 1 << m

	dfs(0, 0)

	return 1, ans, edge
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
