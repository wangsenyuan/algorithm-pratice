package main

import "testing"

func runSample(t *testing.T, n int, x int, A []int, expect int) {
	res := solve(n, x, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 4, 1
	A := []int{2, 3, 5, 4}
	expect := 3
	runSample(t, n, x, A, expect)
}

func TestSample2(t *testing.T) {
	n, x := 5, 6
	A := []int{1, 1, 3, 4, 4}
	expect := 0
	runSample(t, n, x, A, expect)
}

func TestSample3(t *testing.T) {
	n, x := 2, 10
	A := []int{11, 9}
	expect := -1
	runSample(t, n, x, A, expect)
}

func TestSample4(t *testing.T) {
	n, x := 5, 18
	A := []int{81, 324, 218, 413, 324}
	expect := 3
	runSample(t, n, x, A, expect)
}
