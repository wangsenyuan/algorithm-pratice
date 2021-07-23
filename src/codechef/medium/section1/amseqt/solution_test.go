package main

import "testing"

func runSample(t *testing.T, n int, m int, A []int, expect int) {
	res := solve(n, m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 10
	A := []int{4, 4, 4, 4, 4, 4}
	expect := 32
	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 6, 3
	A := []int{4, 4, 4, 4, 4, 4}
	expect := 4
	runSample(t, n, m, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 8
	A := []int{1, 1, 1, 1, 128}
	expect := 0
	runSample(t, n, m, A, expect)
}

func TestSample4(t *testing.T) {
	n, m := 7, 4
	A := []int{4, 5, 6, 7, 8, 9, 10}
	expect := 3
	runSample(t, n, m, A, expect)
}
