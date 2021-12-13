package p2066

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	// if requests[i] success, cause restricition[x] violates
	set := NewSet(n)
	ans := make([]bool, len(requests))
	for i, req := range requests {
		u, v := req[0], req[1]
		u, v = set.Find(u), set.Find(v)
		if u > v {
			u, v = v, u
		}
		ok := true
		for _, res := range restrictions {
			a, b := res[0], res[1]
			a, b = set.Find(a), set.Find(b)
			if a > b {
				a, b = b, a
			}
			if a == u && b == v {
				ok = false
				break
			}
		}
		if ok {
			set.Union(u, v)
		}
		ans[i] = ok
	}
	return ans
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

func (set *Set) Union(a, b int) {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return
	}
	if set.cnt[pa] < set.cnt[pb] {
		pa, pb = pb, pa
	}
	set.cnt[pa] += set.cnt[pb]
	set.arr[pb] = pa
}
