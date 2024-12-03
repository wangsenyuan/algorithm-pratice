package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, a, cities, workers := process(reader)

	if (len(res) == len(workers)) != expect {
		t.Fatalf("Sample expect %t, but got %v %v %v %v", expect, res, a, cities, workers)
	}

	if !expect {
		return
	}

	n := len(a)

	set := NewDSU(n)

	for i := range len(workers) {
		u := res[i][0]
		if u == 0 {
			continue
		}
		set.Union(u-1, res[i][1]-1)
	}

	for i := 0; i < n; i++ {
		j := a[i]
		if set.Find(i) != set.Find(j) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
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
	s := `4 5
2 2 1 2
3 2 2 3
4 2 3 4
2 2 4 5
1 2 3 4 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
2 2 10 20
3 2 2 3
4 2 3 4
2 2 4 5
1 2 3 4 5
`
	expect := false
	runSample(t, s, expect)
}
