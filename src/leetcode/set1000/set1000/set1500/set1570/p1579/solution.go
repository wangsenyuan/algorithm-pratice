package p1579

func maxNumEdgesToRemove(n int, edges [][]int) int {
	uf := NewUFSet(n)
	var res int
	for _, edge := range edges {
		t, u, v := edge[0], edge[1], edge[2]
		u--
		v--
		if t == 3 && !uf.Union(u, v) {
			res++
		}
	}

	tmp := uf.Clone()

	for _, edge := range edges {
		t, u, v := edge[0], edge[1], edge[2]
		u--
		v--
		if t == 1 && !uf.Union(u, v) {
			res++
		}

		if t == 2 && !tmp.Union(u, v) {
			res++
		}
	}
	r := find(uf.set, 0)
	for i := 0; i < n; i++ {
		if find(uf.set, i) != r {
			return -1
		}
	}
	r = find(tmp.set, 0)
	for i := 0; i < n; i++ {
		if find(tmp.set, i) != r {
			return -1
		}
	}
	return res
}

type UF struct {
	set []int
	cnt []int
}

func NewUFSet(n int) UF {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		cnt[i] = 1
	}
	return UF{set, cnt}
}

func find(set []int, x int) int {
	if set[x] != x {
		set[x] = find(set, set[x])
	}
	return set[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := find(uf.set, a)
	pb := find(uf.set, b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.set[pb] = pa
	uf.cnt[pa] += uf.cnt[pb]
	return true
}

func (uf UF) Clone() UF {
	newSet := make([]int, len(uf.set))
	newCnt := make([]int, len(uf.cnt))
	copy(newSet, uf.set)
	copy(newCnt, uf.cnt)
	return UF{newSet, newCnt}
}
