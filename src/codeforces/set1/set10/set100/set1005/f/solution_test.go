package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	n, edges, res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	// len(res) == expect
	set := NewDSU(n)
	check := func(s string) bool {
		set.Reset()
		var w int
		for i, x := range []byte(s) {
			if x == '1' {
				w++
				e := edges[i]
				u, v := e[0]-1, e[1]-1
				set.Union(u, v)
			}
		}
		return set.size == 1 && w == n-1
	}

	cnt := make(map[string]int)
	for _, s := range res {
		cnt[s]++
		if cnt[s] > 1 {
			t.Fatalf("Sample have duplicates result %v", res)
		}
		if !check(s) {
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

func (set *DSU) Reset() {
	for i := 0; i < len(set.arr); i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = len(set.arr)
}

func TestSample1(t *testing.T) {
	s := `4 4 3
1 2
2 3
1 4
4 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 6 3
1 2
2 3
1 4
4 3
2 4
1 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 6 2
1 2
1 3
2 4
2 5
3 4
3 5
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 14 100
2 8
4 7
1 4
9 1
6 9
8 6
10 2
8 4
1 7
6 5
10 9
3 10
6 2
1 3
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 6 166666
1 2
1 3
2 4
2 5
3 4
3 5
`
	expect := 4
	runSample(t, s, expect)
}