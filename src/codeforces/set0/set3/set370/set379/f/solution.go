package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	steps := make([]int, n)
	for i := 0; i < n; i++ {
		steps[i] = readNum(reader)
	}

	res := solve(steps)

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

const inf = 1 << 60

func solve(steps []int) []int32 {
	// 1, 2, 3, 4
	g := NewGraph(2*len(steps)+10, 2*len(steps)+10)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	n := 4

	for _, v := range steps {
		// 每次加两个节点
		g.AddEdge(v-1, n)
		g.AddEdge(v-1, n+1)
		n += 2
	}

	lg := bits.Len(uint(n))

	fa := make([][]int32, n)
	for i := 0; i < n; i++ {
		fa[i] = make([]int32, lg)
	}
	var dfs func(u int32)

	dep := make([]int32, n)

	dfs = func(u int32) {
		for i := 1; i < lg; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			fa[v][0] = u
			dep[v] = dep[u] + 1
			dfs(v)
		}
	}
	dfs(0)

	end := []int32{1, 2}
	var dia int32 = 2

	lca := func(u int32, v int32) int32 {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := lg - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := lg - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	getDistance := func(u int32, v int32) int32 {
		p := lca(u, v)
		return dep[u] + dep[v] - 2*dep[p]
	}

	add := func(u int32) {
		// 将这个节点的位置设置上去
		// may update
		a := getDistance(end[0], u)
		b := getDistance(end[1], u)
		if a > dia {
			dia = a
			end[1] = u
		}
		if b > dia {
			dia = b
			end[0] = u
		}
	}
	var m int32 = 4

	ans := make([]int32, len(steps))

	for i := range steps {
		add(m)
		add(m + 1)
		m += 2
		ans[i] = dia
	}

	return ans
}

type Graph struct {
	nodes []int32
	next  []int32
	to    []int32
	cur   int32
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int32, n)
	next := make([]int32, e)
	to := make([]int32, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = int32(v)
}
