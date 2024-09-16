package main

import "testing"

func runSample(t *testing.T, a, b int, k int, h []int, expect int) {
	res := solve(a, b, k, h)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b, k := 2, 3, 3
	h := []int{7, 10, 50, 12, 1, 8}
	expect := 5
	runSample(t, a, b, k, h, expect)
}

func TestSample2(t *testing.T) {
	a, b, k := 1, 100, 99
	h := []int{100}
	expect := 1
	runSample(t, a, b, k, h, expect)
}
