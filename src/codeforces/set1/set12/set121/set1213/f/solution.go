package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _, _ := process(reader)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res)
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

func process(reader *bufio.Reader) (s string, p []int, q []int, k int) {
	n, k := readTwoNums(reader)
	p = readNNums(reader, n)
	q = readNNums(reader, n)
	s = solve(k, p, q)
	return
}

func solve(k int, p []int, q []int) string {
	n := len(p)

	// 每个节点最多有两条边
	g := NewGraph(n, 2*n)

	add := func(u int, v int) {
		u--
		v--
		g.AddEdge(u, v)
	}

	for i := 0; i+1 < n; i++ {
		add(p[i], p[i+1])
		add(q[i], q[i+1])
	}

	dfn := make([]int, n)
	low := make([]int, n)
	for i := 0; i < n; i++ {
		dfn[i] = -1
	}
	// vis[i] 表示i是否在路径上
	vis := make([]bool, n)
	stack := make([]int, n)
	var top int
	var clock int

	var cid int

	id := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		dfn[u] = clock
		low[u] = clock
		clock++
		vis[u] = true
		stack[top] = u
		top++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dfn[v] < 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}

		if low[u] == dfn[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				id[v] = cid
				if u == v {
					break
				}
			}
			cid++
		}
	}

	for i := 0; i < n; i++ {
		if dfn[i] < 0 {
			dfs(i)
		}
	}

	if cid < k {
		return ""
	}
	// cid >= k
	g2 := NewGraph(cid, 2*n)
	deg := make([]int, cid)

	for i := 0; i+1 < n; i++ {
		u, v := p[i]-1, p[i+1]-1
		if id[u] != id[v] {
			g2.AddEdge(id[u], id[v])
			deg[id[v]]++
		}
		u, v = q[i]-1, q[i+1]-1
		if id[u] != id[v] {
			g2.AddEdge(id[u], id[v])
			deg[id[v]]++
		}
	}

	buf := make([]byte, cid)

	for i := 0; i < cid; i++ {
		buf[i] = ' '
	}

	var dfs2 func(u int, c int)

	dfs2 = func(u int, c int) {
		buf[u] = byte(c + 'a')
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			if buf[v] == ' ' {
				// 有可能有重边，所以这里的检查时必要的
				dfs2(v, min(c+1, 25))
			}
		}
	}
	var c int
	for i := 0; i < cid; i++ {
		if deg[i] == 0 {
			dfs2(i, min(c, 25))
			c++
		}
	}

	ans := make([]byte, n)

	for i := 0; i < n; i++ {
		ans[i] = buf[id[i]]
	}

	return string(ans)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
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
