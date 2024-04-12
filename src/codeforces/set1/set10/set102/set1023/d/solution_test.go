package main

import "testing"

func runSample(t *testing.T, q int, a []int, expect bool) {
	res := solve(q, a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	first := make([]int, q+1)
	for i := 0; i <= q; i++ {
		first[i] = -1
	}
	last := make([]int, q+1)

	n := len(a)

	tr := NewSegTree(n, q+1, min)

	for i := 0; i < n; i++ {
		if a[i] == 0 {
			t.Fatalf("Sample result not correct, value = 0 at position %d", i)
		}
		tr.Update(i, a[i])
		if first[a[i]] < 0 {
			first[a[i]] = i
		}
		last[a[i]] = i
	}

	for i := q; i > 0; i-- {
		if first[i] < 0 {
			continue
		}
		l, r := first[i], last[i]
		u := tr.Get(l, r+1)
		if u < i {
			t.Fatalf("Sample result not correct")
		}
	}
}

func TestSample1(t *testing.T) {
	q := 3
	a := []int{1, 0, 2, 3}
	expect := true
	runSample(t, q, a, expect)
}

func TestSample2(t *testing.T) {
	q := 10
	a := []int{10, 10, 10}
	expect := true
	runSample(t, q, a, expect)
}

func TestSample3(t *testing.T) {
	q := 6
	a := []int{6, 5, 6, 2, 2}
	expect := false
	runSample(t, q, a, expect)
}

func TestSample4(t *testing.T) {
	q := 5
	a := []int{0, 0, 0}
	expect := true
	runSample(t, q, a, expect)
}
