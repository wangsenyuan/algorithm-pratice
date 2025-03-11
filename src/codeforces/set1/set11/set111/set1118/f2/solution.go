package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	g := make([][]int, n)

	for range n - 1 {
		u, v := readTwoNums(reader)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	return solve(k, a, g)
}

const mod = 998244353

func mul(num ...int) int {
	res := 1
	for _, x := range num {
		res *= x
		res %= mod
	}
	return res
}

func add(a, b int) int {
	return (a + b) % mod
}

type pair struct {
	first  int
	second int
}

func solve(k int, color []int, g [][]int) int {
	n := len(color)
	h := bits.Len32(uint32(n))
	fa := make([][]int, n)
	for i := range n {
		fa[i] = make([]int, h)
	}

	dep := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for _, v := range g[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
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

	l := make([]int, k+1)
	for i := range k + 1 {
		l[i] = -1
	}

	for u := range n {
		if color[u] != 0 {
			if l[color[u]] < 0 {
				l[color[u]] = u
			} else {
				l[color[u]] = lca(l[color[u]], u)
			}
		}
	}

	for i := 1; i <= k; i++ {
		// color[l[i]] = 0表示这是一个空白点
		if color[l[i]] != 0 && color[l[i]] != i {
			return 0
		}
		color[l[i]] = i
	}

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 2)
	}

	var dfs2 func(p int, u int) int

	dfs2 = func(p int, u int) int {
		var cur []pair
		sub_color := -1
		for _, v := range g[u] {
			if p != v {
				c := dfs2(u, v)
				if c == -2 {
					// no answer
					return c
				}
				if c != 0 {
					cur = append(cur, pair{c, v})
				}
				if c > 0 {
					if sub_color > 0 && sub_color != c {
						return -2
					}
					sub_color = c
				}
			}
		}

		// 有多个未完成的颜色分组，或者和当前的颜色不一致
		if color[u] != 0 && sub_color > 0 && sub_color != color[u] {
			return -2
		}
		if color[u] == 0 && len(cur) == 0 {
			return 0
		}
		if color[u] == 0 && sub_color < 0 {
			// c 都是-1
			pr := []int{1, 1}
			for _, x := range cur {
				v := x.second
				tmp := mul(pr[len(pr)-1], add(dp[v][0], dp[v][1]))
				pr = append(pr, tmp)
			}
			su := []int{1, 1}

			for i := len(cur) - 1; i >= 0; i-- {
				v := cur[i].second
				tmp := mul(su[len(su)-1], add(dp[v][0], dp[v][1]))
				su = append(su, tmp)
			}
			reverse(su)
			dp[u][1] = 0
			for i, x := range cur {
				dp[u][1] = add(dp[u][1], mul(pr[i], su[i+1], dp[x.second][1]))
			}
			dp[u][0] = pr[len(cur)]
			return -1
		}
		dp[u][1] = 1
		for _, x := range cur {
			if x.first == -1 {
				dp[u][1] = mul(dp[u][1], add(dp[x.second][0], dp[x.second][1]))
			} else {
				dp[u][1] = mul(dp[u][1], dp[x.second][1])
			}
		}
		if color[u] == 0 {
			return sub_color
		}
		// color[u] > 0
		if l[color[u]] == u {
			return -1
		}
		return color[u]
	}

	c := dfs2(-1, 0)
	if c == -2 {
		return 0
	}
	return dp[0][1]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
