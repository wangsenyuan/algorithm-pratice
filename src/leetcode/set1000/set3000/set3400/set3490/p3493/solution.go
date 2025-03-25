package p3493

import "math/bits"

func numberOfComponents(properties [][]int, k int) int {
	type state struct {
		lo uint64
		hi uint64
	}

	create := func(cur []int) state {
		var lo, hi uint64
		for _, p := range cur {
			if p >= 64 {
				hi |= 1 << (p - 64)
			} else {
				lo |= 1 << p
			}
		}
		return state{lo, hi}
	}

	n := len(properties)
	states := make([]state, n)
	for i, cur := range properties {
		states[i] = create(cur)
	}
	set := NewDSU(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cnt := bits.OnesCount64(states[i].lo & states[j].lo)
			cnt += bits.OnesCount64(states[i].hi & states[j].hi)
			if cnt >= k {
				set.Union(i, j)
			}
		}
	}

	return set.sz
}

type DSU struct {
	arr []int
	cnt []int
	sz  int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt, n}
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
	this.sz--
	return true
}
