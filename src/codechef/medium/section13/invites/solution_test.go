package main

import "testing"

func runSample(t *testing.T, n int, enemies []int, expect int) {
	res := solve(n, enemies)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, enemies, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	enemies := []int{2, 3, 4, 1}
	expect := 7
	runSample(t, n, enemies, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	enemies := []int{2, 1, 2}
	expect := 5
	runSample(t, n, enemies, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	enemies := []int{2, 1}
	expect := 3
	runSample(t, n, enemies, expect)
}
