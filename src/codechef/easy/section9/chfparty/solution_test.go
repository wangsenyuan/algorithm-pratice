package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Fatalf("sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []int{0, 0}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, []int{3, 1, 0, 0, 5, 5}, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, []int{1, 2, 3}, 0)
}
