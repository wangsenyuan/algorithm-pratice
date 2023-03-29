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
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}

		ans, res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", ans))
		if ans == 1 {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
		} else {
			buf.WriteString(fmt.Sprintf("%d", len(res)))
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf(" %d", res[i]))
			}
		}
		buf.WriteByte('\n')
		if buf.Len() > 10000 {
			fmt.Print(buf.String())
			buf.Reset()
		}
	}

	if buf.Len() == 0 {
		return
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

func solve(n int, E [][]int) (int, []int) {
	g := NewGraph(n+1, 2*n)

	for _, e := range E {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	color := make([]int, n+1)
	for i := 0; i <= n; i++ {
		color[i] = -1
	}

	var dfs func(u int, c int)

	in_tree := make([]bool, len(E)*2+3)

	dfs = func(u int, c int) {
		color[u] = c
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if color[v] < 0 {
				in_tree[i] = true
				dfs(v, 1^c)
			}
		}
	}

	dfs(1, 1)

	old := make([]int, n+1)

	for i := 0; i <= n; i++ {
		old[i] = color[i]
		color[i] = -1
	}

	var dfs2 func(p int, u int, c int) int

	fa := make([]int, n+1)

	dfs2 = func(p int, u int, c int) int {
		color[u] = c
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if in_tree[i] || in_tree[i^1] || p == v {
				// in the tree
				continue
			}
			fa[v] = u
			if color[v] < 0 {
				w := dfs2(u, v, 1^c)
				if w > 0 {
					return w
				}
			} else if color[v] == c {
				// an odd length cycle
				return v
			}
			// else an even length cycle, ignore it
		}
		return 0
	}

	for u := 1; u <= n; u++ {
		if color[u] < 0 {
			root := dfs2(0, u, 0)
			if root > 0 {
				var res []int

				res = append(res, root)

				for x := fa[root]; x != root; x = fa[x] {
					res = append(res, x)
				}
				return 2, res
			}
		}
	}

	res := make([]int, n+1)

	for u := 1; u <= n; u++ {
		res[u] = ((color[u] << 1) | old[u]) + 1
	}

	return 1, res[1:]
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
	g.next = make([]int, e+10)
	g.to = make([]int, e+10)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
