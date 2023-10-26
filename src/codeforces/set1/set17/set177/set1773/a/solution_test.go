package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	ok, p, q := solve(a)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v %v", expect, ok, p, q)
	}
	if !expect {
		return
	}
	n := len(a)

	for i := 0; i < n; i++ {
		p[i]--
		q[i]--
	}
	for i := 0; i < n; i++ {
		tmp := a[p[q[i]]]
		if tmp != i {
			t.Fatalf("Sample result %v %v, not correct", p, q)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 4, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 1, 4, 2, 3}
	expect := true
	runSample(t, a, expect)
}
