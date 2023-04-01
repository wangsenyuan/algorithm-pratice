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
	A := make([][]int, m)
	for i := 0; i < m; i++ {
		A[i] = readNNums(reader, 3)
	}
	dir, x := solve(n, A)

	if len(dir) == 0 {
		fmt.Println("NO")
	} else {
		var buf bytes.Buffer
		buf.WriteString("YES\n")
		for i := 0; i < n; i++ {
			buf.WriteByte(dir[i])
			buf.WriteString(fmt.Sprintf(" %d\n", x[i]))
		}
		fmt.Print(buf.String())
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, A [][]int) (string, []int) {
	g := NewGraph(n, 2*len(A))
	for _, a := range A {
		_, u, v := a[0], a[1], a[2]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	// 0 for L, 1 for R
	ori := make([]int, n)
	for i := 0; i < n; i++ {
		ori[i] = -1
	}

	var dfs func(u, o int) bool

	dfs = func(u, o int) bool {
		if ori[u] != -1 {
			return ori[u] == o
		}
		ori[u] = o
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, 1^o) {
				return false
			}
		}
		return true
	}

	for u := 0; u < n; u++ {
		if ori[u] < 0 && !dfs(u, 0) {
			return "", nil
		}
	}

	// possible
	g2 := NewGraph(n, len(A))

	for _, a := range A {
		t, u, v := a[0], a[1], a[2]
		u--
		v--

		if ori[u] == 1 {
			u, v = v, u
		}
		// u directs to L, and v directs to R
		if t == 2 {
			u, v = v, u
		}
		// u is before v, v depends on u
		g2.AddEdge(v, u)
	}

	ord := make([]int, 0, n)
	vis := make([]int, n)
	var dfs2 func(u int) bool

	dfs2 = func(u int) bool {
		if vis[u] == 1 {
			// loop back
			return false
		}
		vis[u]++

		for i := g2.node[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			if vis[v] < 2 {
				if !dfs2(v) {
					return false
				}
			}
		}

		vis[u]++

		ord = append(ord, u)

		return true
	}

	for u := 0; u < n; u++ {
		if vis[u] == 0 && !dfs2(u) {
			return "", nil
		}
	}

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		if ori[i] == 0 {
			buf[i] = 'L'
		} else {
			buf[i] = 'R'
		}
	}

	ans := make([]int, n)

	for i, u := range ord {
		ans[u] = i
	}

	return string(buf), ans
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
	g.to[g.cur] = v
	g.node[u] = g.cur
}
