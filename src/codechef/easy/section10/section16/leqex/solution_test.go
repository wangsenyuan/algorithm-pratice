package main

import "testing"

func runSample(t *testing.T, n int, color []int, expect int) {
	res := solve(n, color)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, color, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 14
	color := []int{5, 4, 2, 2, 3, 2, 1, 3, 2, 7, 4, 9, 9, 9}
	expect := 3
	runSample(t, n, color, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	color := []int{1, 2, 1}
	expect := 1
	runSample(t, n, color, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	color := []int{1, 1, 1}
	expect := 1
	runSample(t, n, color, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	color := []int{1, 2, 3, 4, 1}
	expect := 0
	runSample(t, n, color, expect)
}
