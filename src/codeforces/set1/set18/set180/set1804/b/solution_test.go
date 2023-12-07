package main

import "testing"

func runSample(t *testing.T, n int, k int, d int, w int, x []int, expect int) {
	res := solve(n, k, d, w, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, d, w := 6, 3, 5, 3
	x := []int{1, 2, 3, 10, 11, 18}
	expect := 2
	runSample(t, n, k, d, w, x, expect)
}

func TestSample2(t *testing.T) {
	n, k, d, w := 6, 4, 0, 0
	x := []int{3, 3, 3, 3, 3, 4}
	expect := 3
	runSample(t, n, k, d, w, x, expect)
}

func TestSample3(t *testing.T) {
	n, k, d, w := 9, 10, 2, 2
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	expect := 2
	runSample(t, n, k, d, w, x, expect)
}
