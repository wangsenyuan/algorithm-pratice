package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	a, b := solve(n, m)

	if abs(a-b) != expect {
		t.Errorf("Sample expect %d, but got %d %d", expect, a, b)
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 6, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 8, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 89, 77)
}
