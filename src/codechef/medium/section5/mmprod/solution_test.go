package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	A := []int{5, -3, -4, 6}
	expect := 30
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	A := []int{-5, 3, 4, -6}
	expect := 30
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 2
	A := []int{10000, 100000, 100000}
	expect := 999999937
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n, k := 2, 1
	A := []int{-1, -2}
	expect := 1000000006
	runSample(t, n, k, A, expect)
}

func TestSample5(t *testing.T) {
	n, k := 4, 2
	A := []int{-1, -2, 3, -5}
	expect := 10
	runSample(t, n, k, A, expect)
}

func TestSample6(t *testing.T) {
	n, k := 4, 2
	A := []int{-1, 2, -3, 5}
	expect := 10
	runSample(t, n, k, A, expect)
}
