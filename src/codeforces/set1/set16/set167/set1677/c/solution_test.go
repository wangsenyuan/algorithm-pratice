package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := int(solve(A, B))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 4, 3, 2, 6}
	B := []int{5, 3, 1, 4, 6, 2}
	expect := 18
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 5, 4, 6, 2, 1}
	B := []int{3, 6, 4, 5, 2, 1}
	expect := 10
	runSample(t, A, B, expect)
}
