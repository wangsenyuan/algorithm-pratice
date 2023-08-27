package p2836

const H = 20

func getMaxFunctionValue(receiver []int, k int64) int64 {
	n := len(receiver)

	g := NewGraph(n, n*2)

	for i := 0; i < n; i++ {
		j := receiver[i]
		if i != j {
			g.AddEdge(i, j)
			g.AddEdge(j, i)
		}
	}

	in_cycle := make([]int, n)
	pos_cycle := make([]int, n)
	for i := 0; i < n; i++ {
		in_cycle[i] = -1
		pos_cycle[i] = -1
	}

	vis := make([]int, n)

	var dfs func(u int)

	var cycles [][]int

	dfs = func(u int) {
		vis[u]++
		v := receiver[u]
		if vis[v] == 0 {
			dfs(v)
		} else if vis[v] == 1 {
			// an cycle
			if in_cycle[v] < 0 {
				var j int
				x := v
				var cur []int
				for in_cycle[x] < 0 {
					cur = append(cur, x)
					in_cycle[x] = len(cycles)
					pos_cycle[x] = j
					j++
					x = receiver[x]
				}
				cycles = append(cycles, cur)
			}
			//vis[v]++
		}

		vis[u]++
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}

	depth := make([]int, n)
	sum := make([]int64, n)
	P := make([][]int, n)

	var sum_in_cycle [][]int64

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		P[u] = make([]int, H)
		P[u][0] = p

		for i := 1; i < H; i++ {
			if P[u][i-1] == -1 {
				P[u][i] = -1
			} else {
				P[u][i] = P[P[u][i-1]][i-1]
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p || in_cycle[v] >= 0 {
				continue
			}
			depth[v] = depth[u] + 1
			sum[v] = sum[u] + int64(v)
			dfs2(u, v)
		}
	}

	for _, cur := range cycles {
		pref := make([]int64, len(cur))
		for i, u := range cur {
			sum[u] = int64(u)
			dfs2(-1, u)
			pref[i] = int64(u)
			if i > 0 {
				pref[i] += pref[i-1]
			}
		}
		sum_in_cycle = append(sum_in_cycle, pref)
	}

	find_in_cycle := func(u int, id int, k int64) int64 {
		pref := sum_in_cycle[id]

		m := len(pref)
		x, y := k/int64(m), int(k%int64(m))
		ans := int64(u) + x*pref[m-1]
		// 还需要知道u在cycle的哪个位置
		if y > 0 {
			i := pos_cycle[u]
			if m-i > y {
				ans += pref[i+y] - pref[i]
			} else {
				ans += pref[m-1] - pref[i] + pref[y-(m-i)]
			}
		}
		return ans
	}

	kth_ancestor := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = P[u][i]
			}
		}
		return u
	}

	find := func(u int) int64 {
		if in_cycle[u] >= 0 {
			return find_in_cycle(u, in_cycle[u], k)
		}
		if int64(depth[u]) >= k {
			p := kth_ancestor(u, int(k))
			if P[p][0] >= 0 {
				return sum[u] - sum[P[p][0]]
			}
			return sum[u]
		}

		v := kth_ancestor(u, depth[u])
		ans := sum[u] - sum[v]
		return ans + find_in_cycle(v, in_cycle[v], k-int64(depth[u]))
	}

	var best int64

	for i := 0; i < n; i++ {
		best = max(best, find(i))
	}

	return best
}

func max(a, b int64) int64 {
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
	e++
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
