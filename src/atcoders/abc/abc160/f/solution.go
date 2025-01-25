package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(num ...int) int {
	res := 1
	for _, x := range num {
		res = res * x % mod
	}
	return res
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func div(a, b int) int {
	return mul(a, inverse(b))
}

func solve(n int, edges [][]int) []int {
	F := make([]int, n+1)
	I := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[n] = inverse(F[n])
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	f := make([]int, n)
	f[0] = 1

	sz := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
		f[0] = mul(f[0], sz[u])
	}
	dfs(-1, 0)

	// fp 表示在不考虑u子树的情况下， p的计数
	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		if p >= 0 {
			// 这个是p的sz
			f[u] = mul(div(f[p], sz[u]), n-sz[u])
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
			}
		}
	}

	dfs2(-1, 0)

	for i := 0; i < n; i++ {
		f[i] = div(F[n], f[i])
	}

	return f
}

func solve1(n int, edges [][]int) []int {
	F := make([]int, n+1)
	I := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[n] = inverse(F[n])
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], I[r], I[n-r])
	}

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	dp := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
		dp[u] = 1
		tmp := sz[u] - 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dp[u] = mul(dp[u], nCr(tmp, sz[v]), dp[v])
				tmp -= sz[v]
			}
		}
	}
	dfs(-1, 0)

	f := make([]int, n)
	// fp 表示在不考虑u子树的情况下， p的计数
	var dfs2 func(p int, u int, fp int)
	dfs2 = func(p int, u int, fp int) {
		p_sz := n - sz[u]
		if p >= 0 {
			tmp := n - 1
			f[u] = mul(fp, nCr(tmp, p_sz))
			tmp -= p_sz
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					f[u] = mul(f[u], nCr(tmp, sz[v]), dp[v])
					tmp -= sz[v]
				}
			}
		} else {
			f[u] = dp[u]
		}
		// 这里不知道咋算了
		// 因为去掉一个子节点，增加一个子节点的时候，一开始的数量是不一样的
		// 也就是说，一开始是 6 * 4 * 2
		// 但现在有可能变成 7 * 5 * 3
		// 但是其他的是一样的
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				// 排除掉v的计数
				fu := div(dp[u], mul(nCr(sz[u]-1, sz[v]), dp[v]))
				// 其他的子节点 是从 sz[u] - sz[v]开始算的
				// 增加p的计数
				fu = mul(fu, nCr(n-sz[v]-1, p_sz), fp)
				dfs2(u, v, fu)
			}
		}
	}

	dfs2(-1, 0, 1)

	return f
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
