package main

import "testing"

func runSample(t *testing.T, A []int, hd int) {
	dist, C := solve(A)

	if hd != dist {
		t.Errorf("sample: %v, should give %d, but got %d, %v", A, hd, dist, C)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 1}, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{2, 6, 5, 2}, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{1, 2, 3, 4}, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{3, 2, 4, 2}, 4)
}

func TestSample6(t *testing.T) {
	runSample(t, []int{3, 4, 2, 2}, 4)
}
