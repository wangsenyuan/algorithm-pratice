package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{0, 90, 180, 270}
	runSample(t, n, A, 0)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{90, 210}
	runSample(t, n, A, 1)
}
func TestSample3(t *testing.T) {
	n := 2
	A := []int{0, 1}
	runSample(t, n, A, 358)
}
