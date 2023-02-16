package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{7, 9, 3, 14, 63}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{2, 14, 42}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{45, 9, 3, 18}
	expect := 1
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	A := []int{2, 2, 8}
	expect := 0
	runSample(t, n, A, expect)
}
