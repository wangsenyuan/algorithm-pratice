package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := int(solve(A, B))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, -1, -1}
	B := []int{2, -2, 3}
	expect := 3
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, -1}
	B := []int{4, 5}
	expect := 9
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-1, -2, -3}
	B := []int{-2, -3, -1}
	expect := 0
	runSample(t, A, B, expect)
}
