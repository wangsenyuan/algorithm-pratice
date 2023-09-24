package p2864

func countPaths(n int, edges [][]int) int64 {
	var primes []int
	lpf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			lpf[i] = i
		}
		for j := 0; j < len(primes); j++ {
			if primes[j]*i > n {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	g := NewGraph(n+1, 2*(n+1))

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([]int64, n+1)
	// dp[u] 表示u的子树中， 从u出发的每条路径，从第一个质数节点后，到下一个质数间有多少个节点
	fp := make([]int64, n+1)
	// fp[u] 表示u的子树中，从u出发的每条路径到第一个质数节点中间的节点数
	xp := make([]int64, n+1)
	var ans int64

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		dp[u] = 0

		if lpf[u] != u {
			fp[u] = 1
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			dfs(u, v)

			if lpf[u] == u {
				ans += (fp[u] - xp[u]) * (fp[v] - xp[v])
			} else {
				ans += (fp[u] - xp[u]) * dp[v]
				ans += dp[u] * (fp[v] - xp[v])
			}

			dp[u] += dp[v]
			fp[u] += fp[v]
			xp[u] += xp[v]
		}
		if lpf[u] == u {
			dp[u] = fp[u] - xp[u] + 1
			ans += fp[u] - xp[u]
			fp[u] = 1
			xp[u] = 1
		}
	}

	dfs(1, 1)
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
