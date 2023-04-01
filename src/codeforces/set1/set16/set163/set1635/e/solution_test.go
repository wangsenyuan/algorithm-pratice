package main

import "testing"

func runSample(t *testing.T, n int, A [][]int, expect bool) {
	ori, x := solve(n, A)

	if len(ori) == n != expect {
		t.Fatalf("Sample expect %t, but got %s %v", expect, ori, x)
	}
	if !expect {
		return
	}
	// check
	for _, a := range A {
		w, u, v := a[0], a[1], a[2]
		u--
		v--
		if ori[u] == ori[v] {
			t.Fatalf("Sample result %s not correct, as %d %d has same direction", ori, u, v)
		}
		if ori[u] == 'R' {
			u, v = v, u
		}
		//ori[u] == 'L'
		if w == 1 {
			// irr
			if x[u] > x[v] {
				t.Fatalf("Sample result %s %v, not correct", ori, x)
			}
		}
		if w == 2 {
			if x[u] < x[v] {
				t.Fatalf("Sample result %s %v, not correct", ori, x)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := [][]int{
		{1, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
		{2, 4, 1},
	}
	expect := true
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := [][]int{
		{1, 1, 2},
		{1, 2, 3},
		{1, 1, 3},
	}
	expect := false
	runSample(t, n, A, expect)
}
