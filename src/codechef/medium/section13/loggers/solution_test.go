package main

import "testing"

func runSample(t *testing.T, n int, L []int, expect string) {
	res := solve(n, L)

	if res != expect {
		t.Errorf("Sample %v, expect %s, but got %s", L, expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := []int{2}
	runSample(t, len(L), L, "Alice")
}

func TestSample2(t *testing.T) {
	L := []int{3, 4}
	runSample(t, len(L), L, "Alice")
}

func TestSample3(t *testing.T) {
	L := []int{7, 8, 9}
	runSample(t, len(L), L, "Bob")
}
