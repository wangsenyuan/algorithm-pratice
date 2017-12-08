package main

import "testing"

func runSample(t *testing.T, n int, v []int, expect float64) {
	res := solve(n, v)
	if abs(res-expect) > 0.001 {
		t.Errorf("sample %v should get %.3f, but got %.3f", v, expect, res)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func TestSample1(t *testing.T) {
	runSample(t, 1, []int{100}, 100.0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, []int{42, 0}, 21.0)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, []int{1, 4, 9}, 9.5)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, []int{5, 5, 5, 5}, 10.0)
}
