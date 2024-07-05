package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(a, edges)

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

func solve(a []int, edges [][]int) int {
	n := len(a)

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	arr := make([]int, n+1)
	copy(arr[1:], a)

	n++

	// dp[i] 是到i的最大的连续的节点之和
	dp := make([]int, n)
	next := make([][]int, n)
	// fp[u]是子树u中最大的一条完整path的sum
	fp := make([]int, n)

	// 有两种情况，
	// 1. 这两条path在某个节点处相接了，组成一条完整的path
	// 2. 它们没有相接，是两条path。 所以会存在一个点u，一条path在它的某个子树v的内部，一条path经过了u，但是在v的外部
	// 3. 其实还有第三种情况

	var dfs func(p int, u int)
	var best int

	dfs = func(p int, u int) {
		next[u] = []int{0, 0, 0}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if dp[next[u][0]] <= dp[v] {
					next[u][2] = next[u][1]
					next[u][1] = next[u][0]
					next[u][0] = v
				} else if dp[next[u][1]] <= dp[v] {
					next[u][2] = next[u][1]
					next[u][1] = v
				} else if dp[next[u][2]] <= dp[v] {
					next[u][2] = v
				}
				best = max(best, fp[u]+fp[v])
				fp[u] = max(fp[u], fp[v])
			}
		}

		fp[u] = max(fp[u], dp[next[u][0]]+dp[next[u][1]]+arr[u])
		dp[u] = dp[next[u][0]] + arr[u]
	}

	dfs(0, 1)

	var dfs2 func(p int, u int, from int)

	dfs2 = func(p int, u int, from int) {

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				// 一条路径在子树v中, 另外一个经过u
				if v == next[u][0] {
					best = max(best, dp[next[u][1]]+max(from, dp[next[u][2]])+arr[u]+fp[v])
					dfs2(u, v, max(dp[next[u][1]], from)+arr[u])
				} else {
					tmp := dp[next[u][0]] + arr[u]
					if v == next[u][1] {
						tmp += max(from, dp[next[u][2]])
					} else {
						tmp += max(from, dp[next[u][1]])
					}
					best = max(best, tmp+fp[v])
					dfs2(u, v, max(dp[next[u][0]], from)+arr[u])
				}

			}
		}
	}

	dfs2(0, 1, 0)

	return best
}

func solve1(a []int, edges [][]int) int {
	n := len(a)

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	arr := make([]int, n+1)
	copy(arr[1:], a)

	n++

	var ans int

	var dfs func(int, int) (int, int, int)
	dfs = func(v, fa int) (int, int, int) {
		val := arr[v]
		maxChain := val
		maxPathV := val
		maxPathW := 0
		maxChainPath := val
		for i := g.nodes[v]; i > 0; i = g.next[i] {
			w := g.to[i]
			if w == fa {
				continue
			}
			chainW, pathW, chainPathW := dfs(w, v)
			ans = max(ans, maxPathV+pathW, maxPathW+pathW, maxChain+chainPathW, maxChainPath+chainW)
			maxChainPath = max(maxChainPath, chainPathW+val, maxPathW+chainW+val, maxChain+pathW)
			maxPathV = max(maxPathV, maxChain+chainW)
			maxPathW = max(maxPathW, pathW)
			maxChain = max(maxChain, chainW+val)
		}
		return maxChain, max(maxPathV, maxPathW), maxChainPath
	}
	dfs(1, 0)

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
