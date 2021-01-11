package p1721

import "sort"

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	set := NewDSW(n)
	for _, swap := range allowedSwaps {
		a, b := swap[0], swap[1]
		set.Union(a, b)
	}

	mem := make(map[int][]int)

	for i := 0; i < n; i++ {
		pi := set.Find(i)

		if v, found := mem[pi]; !found {
			mem[pi] = make([]int, 0, 10)
			mem[pi] = append(mem[pi], i)
		} else {
			mem[pi] = append(v, i)
		}
	}
	A := make([]int, n)
	B := make([]int, n)

	countDist := func(v []int) int {
		var m int
		for _, i := range v {
			A[m] = source[i]
			B[m] = target[i]
			m++
		}
		sort.Ints(A[:m])
		sort.Ints(B[:m])
		var cnt int
		var i, j int
		for i < m && j < m {
			if A[i] == B[j] {
				cnt++
				i++
				j++
			} else if A[i] < B[j] {
				i++
			} else {
				j++
			}
		}
		return m - cnt
	}
	var res int
	for _, v := range mem {
		res += countDist(v)
	}

	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSW(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (set *DSU) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *DSU) Union(a, b int) bool {
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
