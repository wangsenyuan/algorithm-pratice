package main

import "testing"

func runSample(t *testing.T, w []int, l int, r int, L int, R int, expect int64) {
	res := solve(w, l, r, L, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r, L, R := 4, 4, 19, 1
	w := []int{42, 3, 99}
	var expect int64 = 576
	runSample(t, w, l, r, L, R, expect)
}

func TestSample2(t *testing.T) {
	l, r, L, R := 7, 2, 3, 9
	w := []int{1, 2, 3, 4}
	var expect int64 = 34
	runSample(t, w, l, r, L, R, expect)
}
