package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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
	n, m := readTwoNums(reader)
	w := readNNums(reader, n)
	roads := make([][]int, m)
	for i := range m {
		roads[i] = readNNums(reader, 2)
	}
	s := readNum(reader)
	return solve(n, w, roads, s)
}

const inf = 1 << 30

func solve(n int, w []int, roads [][]int, s int) int {
	g := NewGraph(n, 2*len(roads)+1)
	deg := make([]int, n)
	for _, road := range roads {
		u, v := road[0], road[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}
	dp := make([]int, n)
	vis := make([]bool, n)
	que := make([]int, n)
	var head, tail int

	s--
	for i := 0; i < n; i++ {
		if i != s && deg[i] == 1 {
			// a leaf
			que[head] = i
			head++
			vis[i] = true
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dp[v] = max(dp[v], dp[u]+w[u])
				deg[v]--
				if deg[v] == 1 && v != s {
					que[head] = v
					head++
					vis[v] = true
				}
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		if vis[i] {
			res = max(res, dp[i]+w[i])
		}
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			res += w[i]
		}
	}

	return res
}

func solve1(n int, w []int, roads [][]int, s int) int {
	s--
	g := NewGraph(n, 2*len(roads)+1)

	for _, road := range roads {
		u, v := road[0], road[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dfn := make([]int, n)
	low := make([]int, n)
	var clock int
	stack := make([]int, n)
	var top int
	vis := make([]bool, n)

	type pair struct {
		first  int
		second int
	}

	var comp []pair
	id := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		clock++
		dfn[u] = clock
		low[u] = clock
		vis[u] = true
		stack[top] = u
		top++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if dfn[v] == 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				// back edge
				low[u] = min(low[u], dfn[v])
			}
		}

		if low[u] == dfn[u] {
			var sum int
			var cnt int
			for top > 0 {
				v := stack[top-1]
				top--
				id[v] = len(comp)
				sum += w[v]
				cnt++
				vis[v] = false
				if u == v {
					break
				}
			}
			comp = append(comp, pair{sum, cnt})
		}
	}

	dfs(-1, s)
	k := len(comp)

	dsu := NewDSU(k)

	g2 := NewGraph(k, 2*len(roads))

	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		if id[u] != id[v] && dsu.Union(id[u], id[v]) {
			// 所属ssc编号越大，越在上面
			if id[u] < id[v] {
				u, v = v, u
			}
			g2.AddEdge(id[u], id[v])
		}
	}

	dp := make([][]int, k)
	// dp[i][0] 表示，进入i后，仍然返回到i的最优解
	// dp[i][1]表示，进入i后，就不返回了（停在某个叶子节点处）

	var dfs2 func(u int)

	dfs2 = func(u int) {
		// 有可能进来但是出不去，所以dp[u][0] = -inf
		dp[u] = []int{-inf, comp[u].first}

		if g2.nodes[u] == 0 {
			// 这是一个叶子节点
			if comp[u].second > 1 {
				// 至少有3个节点，可以返回
				dp[u][0] = comp[u].first
			}
			return
		}

		ok := comp[u].second > 1

		var sum, best int

		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			dfs2(v)

			if dp[v][0] < 0 {
				// 进入节点v，但是无法出来
				best = max(best, dp[v][1])
			} else {
				// 进入节点v后，还能回到节点v，所以也就能回到节点u
				ok = true
				sum += dp[v][0]
				// 如果进入v节点后，不回来，那么它的贡献应该比较大才行
				best = max(best, dp[v][1]-dp[v][0])
			}
		}

		sum += comp[u].first
		dp[u][1] = sum + best

		if ok {
			dp[u][0] = sum
		}
		// else dp[u][0] = -inf
	}

	dfs2(id[s])

	return max(dp[id[s]][0], dp[id[s]][1])
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

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
