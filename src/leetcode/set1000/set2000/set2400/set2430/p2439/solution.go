package p2439

import "sort"

func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	// n <= 20000
	// nums[i] <= 50
	// n * 50 = 1e6
	var sum int
	var mx int
	mn := 1 << 30
	for _, num := range nums {
		sum += num
		mx = max(mx, num)
		mn = min(mn, num)
	}

	if mn == mx {
		// remove all
		return n - 1
	}

	g := NewGraph(n, 2*(n-1))

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	var factors []int
	for i := 2; i <= sum/i; i++ {
		if sum%i == 0 {
			factors = append(factors, i)
			j := sum / i
			if j != i {
				factors = append(factors, j)
			}
		}
	}

	var groups int

	var dfs func(p int, u int, want int) int

	dfs = func(p int, u int, want int) int {
		var sum int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v, want)
				if tmp < 0 {
					return -1
				}
				sum += tmp
			}
		}
		sum += nums[u]

		if sum == want {
			groups++
			return 0
		}
		if sum > want {
			return -1
		}
		return sum
	}

	sort.Ints(factors)
	var ans int

	for _, x := range factors {
		groups = 0
		y := dfs(-1, 0, x)
		if y < 0 {
			continue
		}
		ans = max(ans, groups-1)
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
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
