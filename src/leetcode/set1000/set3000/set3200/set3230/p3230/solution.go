package p3230

import "sort"

func canReachCorner(X int, Y int, circles [][]int) bool {
	sort.Slice(circles, func(i, j int) bool {
		return circles[i][2] < circles[j][2]
	})

	n := len(circles)
	// n for left, n + 1 for bottom, n + 2 for right, n + 3 for top
	set := NewUFSet(n + 4)

	check := func(i, j int) bool {
		dx := circles[j][0] - circles[i][0]
		dy := circles[j][1] - circles[i][1]
		r := circles[j][2] + circles[i][2]
		return dx*dx+dy*dy <= r*r
	}

	for i := 0; i < n; i++ {
		var cnt int
		for j := i + 1; j < n && cnt < 4; j++ {
			if check(i, j) {
				set.Union(i, j)
				cnt++
			}
		}
		if circles[i][0]-circles[i][2] <= 0 {
			set.Union(n, i)
		}
		if circles[i][1]-circles[i][2] <= 0 {
			set.Union(n+1, i)
		}
		if circles[i][0]+circles[i][2] >= X {
			set.Union(n+2, i)
		}
		if circles[i][1]+circles[i][2] >= Y {
			set.Union(n+3, i)
		}
	}
	if set.Find(n) == set.Find(n+1) {
		return false
	}

	if set.Find(n) == set.Find(n+2) {
		return false
	}

	if set.Find(n+1) == set.Find(n+3) {
		return false
	}
	if set.Find(n+2) == set.Find(n+3) {
		return false
	}
	return true
}

type UFSet struct {
	arr  []int
	cnt  []int
	size int
}

func NewUFSet(n int) UFSet {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UFSet{arr, cnt, n}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UFSet) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.cnt[pa] += uf.cnt[pb]
	uf.arr[pb] = pa
	uf.size--
	return true
}
