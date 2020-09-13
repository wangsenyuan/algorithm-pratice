package p1584

import "sort"

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	edges := make([]Edge, 0, n*n/2+1)

	for i := 0; i < n; i++ {
		cur := points[i]
		for j := i + 1; j < n; j++ {
			w := distance(cur, points[j])
			edges = append(edges, Edge{i, j, w})
		}
	}

	sort.Sort(Edges(edges))

	uf := NewUFSet(n)
	var res int
	for _, edge := range edges {
		u, v, w := edge.u, edge.v, edge.w
		if uf.Union(u, v) {
			res += w
		}
	}

	return res
}

func distance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

type Edge struct {
	u, v, w int
}

type Edges []Edge

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].w < edges[j].w
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type UF struct {
	set  []int
	size []int
}

func NewUFSet(n int) *UF {
	set := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		size[i] = 1
	}
	return &UF{set, size}
}

func (this *UF) Find(x int) int {
	set := this.set
	var loop func(x int) int
	loop = func(x int) int {
		if x != set[x] {
			set[x] = loop(set[x])
		}
		return set[x]
	}
	return loop(x)
}

func (this *UF) Union(a, b int) bool {
	pa := this.Find(a)
	pb := this.Find(b)
	if pa == pb {
		return false
	}
	size := this.size
	if size[pa] >= size[pb] {
		this.set[pb] = pa
		size[pa] += size[pb]
	} else {
		this.set[pa] = pb
		size[pb] += size[pa]
	}
	return true
}

func (this UF) Size(x int) int {
	return this.size[x]
}
