package p1916

const MOD = 1000000007

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	g := NewGraph(n, n-1)

	for i := 1; i < n; i++ {
		g.AddEdge(prevRoom[i], i)
	}

	F := make([]int64, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}

	I := make([]int64, n+1)
	I[n] = pow(F[n], MOD-2)

	for i := n - 1; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}

	nCr := func(n, r int) int64 {
		res := F[n]
		res = res * I[r] % MOD
		res = res * I[n-r] % MOD
		return res
	}

	dp := make([]int64, n)
	cnt := make([]int, n)
	var dfs func(u int)

	dfs = func(u int) {
		dp[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			cnt[u] += cnt[v]
			dp[u] = modMul(dp[u], modMul(dp[v], nCr(cnt[u], cnt[v])))
		}
		cnt[u]++
	}

	dfs(0)

	return int(dp[0])
}

func modMul(a, b int64) int64 {
	a *= b
	return a % MOD
}

func pow(a, b int64) int64 {
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = R * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return R
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
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
