package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}

	ask := func(u int, v int) int {
		fmt.Printf("? %d %d\n", u, v)
		return readNum(reader)
	}

	res := solve(n, edges, ask)

	fmt.Printf("! %d\n", res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, edges [][]int, ask func(int, int) int) int {
	g := NewGraph(n+1, 2*n)

	deg := make([]int, n+1)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	leaves := make(map[int]bool)

	for i := 1; i <= n; i++ {
		if deg[i] == 1 {
			leaves[i] = true
		}
	}

	lca, kth := g.prepare(1)

	marked := make([]bool, n+1)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		if marked[u] {
			return
		}
		if leaves[u] {
			delete(leaves, u)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
	}

	mark := func(u int, v int) {
		p := lca(u, v)
		var x int
		if p == u {
			x = kth(v, g.dep[v]-g.dep[p]-1)
		} else {
			x = g.fa[u][0]
		}
		dfs(u, x)
		marked[x] = true
	}

	// len(leaves) >= 2

	for len(leaves) > 2 {
		u, v := getTwo(leaves)

		p := ask(u, v)

		if u == p || v == p {
			return p
		}
		// p 肯定在u, v的路径上， 至少删掉了两个点
		deg[p] -= 2
		if deg[p] == 0 {
			// 很幸运，直接找到了
			return p
		}
		// 把包含u, v 的分支删除掉
		mark(p, u)
		mark(p, v)

		if deg[p] == 1 {
			leaves[p] = true
		}
	}
	// 只剩下一个节点了
	if len(leaves) == 1 {
		return getOne(leaves)
	}
	// 剩下一条线
	u, v := getTwo(leaves)

	return ask(u, v)
}

func getOne(m map[int]bool) int {
	for k := range m {
		return k
	}
	return -1
}

func getTwo(m map[int]bool) (int, int) {
	first, second := -1, -1
	for k := range m {
		if first < 0 {
			first = k
		} else {
			second = k
			break
		}
	}
	return first, second
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	dep   []int
	fa    [][]int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.dep = make([]int, n)
	g.fa = make([][]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)

	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) prepare(root int) (lca func(int, int) int, kth func(int, int) int) {
	n := len(g.nodes)

	h := bits.Len(uint(n))

	dep := g.dep
	fa := g.fa

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(root, root)

	lca = func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	kth = func(u int, k int) int {
		for i := h - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = fa[u][i]
			}
		}
		return u
	}

	return
}
