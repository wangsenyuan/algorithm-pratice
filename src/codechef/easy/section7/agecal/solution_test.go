package main

import "testing"

func runSample(t *testing.T, n int, A []int, yb, mb, db int, yc, mc, dc int, expect int) {
	res := solve(n, A, yb, mb, db, yc, mc, dc)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	yb, mb, db := 2, 1, 1
	yc, mc, dc := 3, 5, 1
	expect := 26
	runSample(t, n, A, yb, mb, db, yc, mc, dc, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{3, 3}
	yb, mb, db := 1, 1, 1
	yc, mc, dc := 2, 2, 2
	expect := 11
	runSample(t, n, A, yb, mb, db, yc, mc, dc, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{1, 1}
	yb, mb, db := 3, 1, 1
	yc, mc, dc := 3, 2, 1
	expect := 2
	runSample(t, n, A, yb, mb, db, yc, mc, dc, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{1, 4, 3, 4, 6}
	yb, mb, db := 3, 5, 6
	yc, mc, dc := 10, 1, 1
	expect := 112
	runSample(t, n, A, yb, mb, db, yc, mc, dc, expect)
}
