package p2972

import "sort"

func placedCoins(edges [][]int, cost []int) []int64 {
	n := len(cost)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	verts := make([]*Vertex, n)

	ans := make([]int64, n)

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		verts[u] = new(Vertex)
		if cost[u] > 0 {
			verts[u].pos = append(verts[u].pos, cost[u])
		} else if cost[u] < 0 {
			verts[u].neg = append(verts[u].neg, cost[u])
		}
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sz += dfs(u, v)
				verts[u].pos = mergePositive(verts[u].pos, verts[v].pos)
				verts[u].neg = mergeNegative(verts[u].neg, verts[v].neg)
			}
		}
		if sz < 3 {
			ans[u] = 1
		} else {
			ans[u] = calcValue(verts[u])
		}
		return sz
	}

	dfs(-1, 0)

	return ans
}

type Vertex struct {
	id  int
	pos []int
	neg []int
}

func calcValue(v *Vertex) int64 {
	var ans int64
	if len(v.pos) == 3 {
		// could have positive value
		ans = int64(v.pos[0]) * int64(v.pos[1]) * int64(v.pos[2])
	}
	if len(v.neg) == 2 && len(v.pos) > 0 {
		ans = max(ans, int64(v.neg[0])*int64(v.neg[1])*int64(v.pos[0]))
	}

	return ans
}

func mergeNegative(a []int, b []int) []int {
	buf := concatAndSort(a, b)
	if len(buf) <= 2 {
		return buf
	}
	return buf[:2]
}

func mergePositive(a []int, b []int) []int {
	buf := concatAndSort(a, b)
	n := len(buf)
	reverse(buf)
	if n <= 3 {
		return buf
	}
	return buf[:3]
}

func concatAndSort(a []int, b []int) []int {
	n := len(a) + len(b)
	buf := make([]int, n)
	copy(buf, a)
	copy(buf[len(a):], b)
	sort.Ints(buf)
	return buf
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
