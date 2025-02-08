package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, expect bool) {
	res := solve(n, a, b)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	set := NewDSU(n)
	set2 := NewDSU(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if res[i][j] == '1' {
				set.Union(i, j)
			} else {
				set2.Union(i, j)
			}
		}
	}

	if set.size != a || set2.size != b {
		t.Fatalf("Sample result %v, not correct", res)
	}
	// 现在检查有几个component
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}

func TestSample1(t *testing.T) {
	n, a, b := 3, 1, 2
	expect := true
	runSample(t, n, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, a, b := 3, 3, 3
	expect := false
	runSample(t, n, a, b, expect)
}

func TestSample3(t *testing.T) {
	n, a, b := 10, 1, 7
	expect := true
	runSample(t, n, a, b, expect)
}

func TestSample4(t *testing.T) {
	n, a, b := 10, 1, 1
	expect := true
	runSample(t, n, a, b, expect)
}

func TestSample5(t *testing.T) {
	n, a, b := 99, 1, 99
	expect := true
	runSample(t, n, a, b, expect)
}
