package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res, _, _ := process(reader)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))

	for _, op := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", op[0], op[1], op[2]))
	}

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (ans [][]int, n int, edges [][]int) {
	n = readNum(reader)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 3)
	}
	ans = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	deg := make([]int, n+1)
	g := NewGraph(n+1, n*2)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
		deg[u]++
		deg[v]++
	}

	for i := 1; i <= n; i++ {
		if deg[i] == 2 {
			return nil
		}
	}

	root := 1
	for deg[root] > 1 {
		root++
	}

	fa := make([]int, n+1)
	fa[root] = -1

	leaves := make([][]int, n+1)

	var dfs1 func(u int)
	dfs1 = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if fa[v] == 0 {
				fa[v] = u
				dfs1(v)
				leaves[u] = append(leaves[u], leaves[v][0])
			}
		}
		if len(leaves[u]) == 0 {
			// a real leaf node
			leaves[u] = append(leaves[u], u)
		}
	}

	dfs1(root)

	var ans [][]int

	addPath := func(v int, x int) {
		if len(leaves[v]) == 1 {
			// 它就是一个leaf node
			ans = append(ans, []int{root, v, x})
			return
		}
		ans = append(ans, []int{root, leaves[v][0], x / 2})
		ans = append(ans, []int{root, leaves[v][1], x / 2})
		ans = append(ans, []int{leaves[v][0], leaves[v][1], -x / 2})
	}

	addEdge := func(v int, x int) {
		addPath(v, x)
		if fa[v] == root {
			return
		}
		addPath(fa[v], -x)
	}

	vis := make([]bool, n+1)

	var dfs2 func(u int)

	dfs2 = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				addEdge(v, g.val[i])
				dfs2(v)
			}
		}
	}

	dfs2(root)

	return ans
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
