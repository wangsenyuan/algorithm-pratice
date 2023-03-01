package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)
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

const H = 20

const INF = 1 << 30

func solve(n int, E [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, 1)
		g.AddEdge(v, u, 0)
	}

	init_inf := func(fp [][]int) {
		for j := 0; j < len(fp); j++ {
			for k := 0; k < len(fp[j]); k++ {
				fp[j][k] = INF
			}
		}
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 5)
		}
		init_inf(dp[i])
	}

	var dfs func(p int, u int)

	fp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		fp[i] = make([]int, 5)
	}

	dfs = func(p int, u int) {
		dp[u][1][1] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				init_inf(fp)
				for a := 0; a <= 2; a++ {
					for b := 0; b <= 2-a; b++ {
						for c := 0; c <= 1; c++ {
							for d := 0; d <= 1; d++ {
								fp[c][a+b] = min(fp[c][a+b], dp[u][c][a]+dp[v][d][b+d]+(1-g.val[i]))
								fp[0][a+b] = min(fp[0][a+b], dp[u][c][a+c]+dp[v][d][b]+g.val[i])
							}
						}
					}
				}
				for j := 0; j < 2; j++ {
					copy(dp[u][j], fp[j])
				}
			}
		}
	}

	dfs(0, 0)

	res := n
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 2; j++ {
			res = min(res, dp[0][i][j])
		}
	}

	return res
}
func solve1(n int, E [][]int) int {
	// n <= 3000
	// 最多从两点出发，可以到达所有的节点
	// 可以用min cost max flow 吗？
	// 每个节点都分成in-out,
	// out[u]到in[v] 如果存在一条正向边, cost 是0， 同时增加一条反向边 out[v]到in[u], cost为1
	// 但似乎没法处理从两个点出发
	// ans[u][v] = 以u, v为root，visit所有节点的翻转数量
	// res =min(ans[u][v])
	// ans[u][v]怎么算?
	// fp[u] 在u的子树中背离u的的边的数量
	// 假设u,v的lca是x，那么 a = fp[0] - fp[x] = 背离0节点，但是不在子树u中的边
	// b = sz[0] - 1 - (sz[x] - 1) 等于除去子树x中的边的数量 b - a = 那些指向节点0的边
	// 这些边是需要被翻转的，还需要处理从路径 [0...x]之间指向0的边，这些是不需要翻转
	// 还需要处理路径[u...x] 和 [v...x]之间的边
	// dp[u][v] = fp[u] + fp[v] + sum()
	// fp[x] = 从x点开始，（不包括子树）需要翻转的边的数量
	// dp[0] = fp[0], dp[x] = 它的子树中需要翻转的数量
	// fp[x] = dp[0] - dp[x] + [0到x]指向x的节点数
	// ans[u][v] = dp[u] + dp[v] + [x..u]指向u的节点数 + [x..v]指向v的节点数 + fp[x]
	// 基本可以处理了
	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, 1)
		g.AddEdge(v, u, -1)
	}
	// dp[u]表示u的子树，中某个点可以访问到所有子树的最小改变数量
	dp := make([]int, n)
	// fp[u]表示从节点u开始访问所有子节点的最小改变数量
	fp := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fp[u] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				fp[u] += fp[v]
				if g.val[i] == -1 {
					// 边是从v到u的
					fp[u]++
				}
			}
		}
		// 从节点u出发
		dp[u] = fp[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				// 从v的某个子节点，reach u， 然后访问其他节点
				tmp := fp[u] - fp[v] + dp[v]
				if g.val[i] == -1 {
					// fp[u]包含了反向边
					tmp--
				}
				if g.val[i] == 1 {
					tmp++
				}
				dp[u] = min(dp[u], tmp)
			}
		}

	}
	arr := make([]int, n)
	check := func(x int) int {
		// 以x为根
		dfs(x, x)

		// 从x的child中，找到两个子树dp[a] + dp[b] + fp[x] - fp[a] - fp[b] - () 最小
		// fp[x] + dp[a] - fp[a] - da + dp[b] - fp[b] + db 最小
		first := -1
		second := -1
		for i := g.nodes[x]; i > 0; i = g.next[i] {
			y := g.to[i]
			tmp := dp[y] - fp[y]
			if g.val[i] == -1 {
				tmp--
			}
			if g.val[i] == 1 {
				tmp++
			}
			arr[y] = tmp
			if first < 0 || arr[first] >= tmp {
				second = first
				first = y
			} else if second < 0 || arr[second] >= tmp {
				second = y
			}
		}
		res := fp[x]
		if first >= 0 {
			res += arr[first]
		}
		if second >= 0 {
			res += arr[second]
		}
		return res
	}

	res := INF
	for i := 0; i < n; i++ {
		res = min(res, check(i))
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, x int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = x
}
