package main

import "testing"

func runSample(t *testing.T, w []int, expect int) {
	res := solve(w)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	w := []int{1, 1, 2, 3, 3}
	expect := 2
	runSample(t, w, expect)
}

func TestSample2(t *testing.T) {
	w := []int{0, 1, 2, 3}
	expect := 4
	runSample(t, w, expect)
}

func TestSample3(t *testing.T) {
	w := []int{688743, 688743, 1975, 688743, 688743, 688743, 688743, 688743, 688743, 0, 0, 688743, 688743}
	expect := 4
	runSample(t, w, expect)
}
