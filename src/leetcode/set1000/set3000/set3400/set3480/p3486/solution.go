package p3486

import "sort"

func longestSpecialPath(edges [][]int, nums []int) []int {
	n := len(nums)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
		g.AddEdge(e[1], e[0], e[2])
	}

	a := sortAndUnique(nums)
	m := len(a)
	pos := make([]int, m)
	for i := range m {
		pos[i] = -1
	}

	dep := make([]int, n)
	sum := make([]int, n)

	ans := make([]int, 2)

	arr := make([]int, n)

	var dfs func(p int, u int, prev []int)
	dfs = func(p int, u int, prev []int) {
		old := make([]int, 2)
		copy(old, prev)
		j := sort.SearchInts(a, nums[u])

		if pos[j] >= 0 {
			c := pos[j]
			// dep[p1] <= dep[p0] 要保证这个成立
			if prev[0] < 0 || dep[c] >= dep[prev[0]] {
				prev[1] = prev[0]
				prev[0] = c
			} else if prev[1] < 0 || dep[c] >= dep[prev[1]] {
				prev[1] = c
			}
		}
		arr[dep[u]] = u
		// 不能到达x
		x := prev[1]
		if x < 0 {
			x = 0
		} else {
			x = arr[dep[x]+1]
		}
		if sum[u]-sum[x] > ans[0] || sum[u]-sum[x] == ans[0] && dep[u]-dep[x] < ans[1] {
			ans[0] = sum[u] - sum[x]
			ans[1] = dep[u] - dep[x]
		}

		old_pos := pos[j]
		pos[j] = u

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sum[v] = sum[u] + g.val[i]
				dep[v] = dep[u] + 1
				dfs(u, v, prev)
			}
		}
		pos[j] = old_pos
		copy(prev, old)
	}

	dfs(-1, 0, []int{-1, -1})
	ans[1]++
	return ans
}

func sortAndUnique(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	var n int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i-1] < res[i] {
			res[n] = res[i-1]
			n++
		}
	}
	return res[:n]
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
