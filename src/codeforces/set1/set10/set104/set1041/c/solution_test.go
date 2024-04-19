package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, m int, d int, a []int, expect int) {
	days, res := solve(m, d, a)

	if days != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, days)
	}
	n := len(a)

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		// 按照天排，在同一天的，按照时间排
		return res[id[i]] < res[id[j]] || res[id[i]] == res[id[j]] && a[id[i]] < a[id[j]]
	})

	for i := 0; i < n; {
		j := i
		for i < n && res[id[i]] == res[id[j]] {
			// 同一天的
			if i > j && a[id[i]]-a[id[i-1]] <= d {
				t.Fatalf("Sample result %v not correct", res)
			}
			i++
		}
	}
}

func TestSample1(t *testing.T) {
	m, d := 5, 3
	a := []int{3, 5, 1, 2}
	expect := 3
	runSample(t, m, d, a, expect)
}

func TestSample2(t *testing.T) {
	m, d := 10, 1
	a := []int{10, 5, 7, 4, 6, 3, 2, 1, 9, 8}
	expect := 2
	runSample(t, m, d, a, expect)
}
