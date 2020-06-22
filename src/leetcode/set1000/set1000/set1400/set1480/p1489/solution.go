package p1489

import (
	"math"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

	E := make([]Edge, len(edges))

	for i := 0; i < len(edges); i++ {
		E[i] = Edge{edges[i][0], edges[i][1], edges[i][2], i}
	}

	sort.Sort(Edges(E))

	uf := CreateUFSet(n)

	getCritical := func(ii int) int {

		var total int

		for i := 0; i < len(E); i++ {
			if i == ii {
				continue
			}

			if uf.Union(E[i].u, E[i].v) {
				total += E[i].c
			}
		}

		if uf.size != 1 {
			return math.MaxInt32
		}

		return total
	}

	tot := getCritical(-1)

	if uf.size != 1 {
		return [][]int{nil, nil}
	}
	uf.Reset()

	critical := make([]bool, len(edges))
	ans1 := make([]int, 0, len(edges))
	for i := 0; i < len(E); i++ {
		tmp := getCritical(i)
		if tmp > tot {
			critical[E[i].i] = true
			ans1 = append(ans1, E[i].i)
		}
		uf.Reset()
	}

	sort.Ints(ans1)

	ans2 := make([]int, 0, len(edges))

	for i := 0; i < len(edges); i++ {
		if critical[i] {
			continue
		}
		u, v, c := edges[i][0], edges[i][1], edges[i][2]
		uf.Union(u, v)
		c += getCritical(-1)

		if c == tot {
			ans2 = append(ans2, i)
		}

		uf.Reset()
	}

	return [][]int{ans1, ans2}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Edge struct {
	u, v, c, i int
}

type Edges []Edge

func (this Edges) Len() int {
	return len(this)
}

func (this Edges) Less(i, j int) bool {
	return this[i].c < this[j].c
}

func (this Edges) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type UFSet struct {
	set  []int
	cnt  []int
	size int
}

func (uf *UFSet) Reset() {
	for i := 0; i < len(uf.set); i++ {
		uf.set[i] = i
		uf.cnt[i] = 1
	}
	uf.size = len(uf.set)
}

func CreateUFSet(n int) UFSet {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = 1
		set[i] = i
	}
	return UFSet{set, cnt, n}
}

func (uf *UFSet) Union(a, b int) bool {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}

	pa := find(a)
	pb := find(b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		uf.set[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	} else {
		uf.set[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	}
	uf.size--
	return true
}
