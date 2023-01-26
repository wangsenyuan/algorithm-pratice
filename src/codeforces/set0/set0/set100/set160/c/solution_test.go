package main

import "testing"

func runSample(t *testing.T, A []int, k int, a, b int) {
	c, d := solve(A, int64(k))
	if a != c || b != d {
		t.Errorf("Sample expect (%d, %d), but got (%d, %d)", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1}
	k := 4
	a, b := 2, 2
	runSample(t, A, k, a, b)
}

func TestSample2(t *testing.T) {
	A := []int{2, 1}
	k := 1
	a, b := 1, 1
	runSample(t, A, k, a, b)
}

func TestSample3(t *testing.T) {
	A := []int{2, 1}
	k := 2
	a, b := 1, 2
	runSample(t, A, k, a, b)
}

func TestSample4(t *testing.T) {
	A := []int{2, 1}
	k := 3
	a, b := 2, 1
	runSample(t, A, k, a, b)
}

func TestSample5(t *testing.T) {
	A := []int{3, 1, 5}
	k := 2
	a, b := 1, 3
	runSample(t, A, k, a, b)
}

func TestSample6(t *testing.T) {
	A := []int{1, 1, 2}
	k := 3
	a, b := 1, 1
	runSample(t, A, k, a, b)
}
