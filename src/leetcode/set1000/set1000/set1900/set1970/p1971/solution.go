package p1971

func validPath(n int, edges [][]int, start int, end int) bool {
	set := NewSet(n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		set.Union(u, v)
	}

	return set.Find(start) == set.Find(end)
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	set := new(Set)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	return set
}

func (set *Set) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *Set) Union(a, b int) bool {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return false
	}
	if set.cnt[pa] < set.cnt[pb] {
		pa, pb = pb, pa
	}
	set.cnt[pa] += set.cnt[pb]
	set.arr[pb] = pa
	return true
}
