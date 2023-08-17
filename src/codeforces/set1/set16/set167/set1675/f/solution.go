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
		//readString(reader)
		n, k := readTwoNums(reader)
		x, y := readTwoNums(reader)
		A := readNNums(reader, k)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, x, y, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
	if bs[0] == '\n' || bs[0] == '\r' {
		return readNNums(reader, n)
	}
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, E [][]int, x int, y int, A []int) int64 {
	// 考虑几种情况,
	// 那些在x一边（远离y)的点, 访问完后，需要回到x
	// 那些在y一边（远离x)的点，需要先到达y，然后访问，再回到y
	// 处于它们中间的部分，从x出发，然后到达y即可
	// 选择x做为root，以y为root的子树中的节点，适用第二种情况
	// 在x的不包含y的子节点中的节点，适用第一种情况
	// x到y过程中的其他部分，适用第三种情况
	// 第一、二中情况，很好处理
	// 第三种情况如何处理呢？也好处理的，只需要已它们中间的节点做为跟节点即可
	x--
	y--

	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	flag := make([]int, n)
	for _, a := range A {
		flag[a-1] = 1
	}

	next := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = -1
	}

	parent := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		parent[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if next[v] != -1 {
					next[u] = v
				}
				flag[u] += flag[v]
			}
		}
		if u == y {
			next[u] = u
		}
	}
	dfs(-1, x)

	var dfs2 func(p int, u int) int64

	dfs2 = func(p int, u int) int64 {
		var res int64

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && flag[v] > 0 {
				res += 2 + dfs2(u, v)
			}
		}

		return res
	}

	var ans int64

	for x != y {
		for i := g.nodes[x]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == next[x] || v == parent[x] {
				continue
			}
			if flag[v] > 0 {
				ans += 2 + dfs2(x, v)
			}
		}
		ans++

		x = next[x]
	}

	ans += dfs2(parent[y], y)

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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
