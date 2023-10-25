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
		n := readNum(reader)
		s := readString(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
		tc--
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

const H = 12

func solve(n int, E [][]int, s string) int {
	// n * n
	// 如果对于x,有两个child, u, v其中都有子串 s, 那么答案至少是 2 * len(s) + 1
	// dp[u][v] = dp[x][y] + 2 if s[u]= s[v] else ...
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	depth := make([]int, n)
	parent := make([][]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {

		parent[u] = make([]int, H)
		parent[u][0] = p
		for i := 1; i < H; i++ {
			parent[u][i] = parent[parent[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				depth[v] = depth[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}

		for i := H - 1; i >= 0; i-- {
			if (depth[u] - (1 << i)) >= depth[v] {
				u = parent[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if parent[u][i] != parent[v][i] {
				u = parent[u][i]
				v = parent[v][i]
			}
		}
		return parent[u][0]
	}

	kthAncestor := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = parent[u][i]
			}
		}
		return u
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// dp[i][j] = max result palindrome between them
	var fn func(u int, v int) int
	fn = func(u int, v int) int {
		if dp[u][v] > 0 {
			return dp[u][v]
		}

		p := lca(u, v)

		if s[u] == s[v] {
			// +2 part
			if u == p || v == p {
				w := u ^ v ^ p
				// direct parent
				if parent[w][0] == p {
					dp[u][v] = 2
				} else {
					x := kthAncestor(w, depth[w]-depth[p]-1)
					y := parent[w][0]
					dp[u][v] = 2 + fn(y, x)
				}
			} else {
				dp[u][v] = 2 + fn(parent[u][0], parent[v][0])
			}
		} else {
			// s[u] != s[v]
			if u == p || v == p {
				// still, one of them is ancestor
				w := u ^ v ^ p
				if parent[w][0] == p {
					dp[u][v] = 1
				} else {
					x := kthAncestor(w, depth[w]-depth[p]-1)
					y := parent[w][0]
					dp[u][v] = max(fn(w, x), fn(y, p))
				}
			} else {
				dp[u][v] = max(fn(u, parent[v][0]), fn(parent[u][0], v))
			}
		}
		return dp[u][v]
	}

	// 还有个可优化的点，只考虑叶子节点对，即可
	var ans int

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ans = max(ans, fn(i, j))
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
