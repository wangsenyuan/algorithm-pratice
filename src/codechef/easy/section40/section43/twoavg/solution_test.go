package main

import "testing"

func runSample(t *testing.T, k int, A []int, B []int, expect int) {
	res := solve(k, A, B)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 9
	A := []int{3, 7, 3, 5, 2, 4}
	B := []int{8, 3, 5}
	expect := 2
	runSample(t, k, A, B, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	A := []int{1, 1}
	B := []int{1, 1}
	expect := -1
	runSample(t, k, A, B, expect)
}

func TestSample3(t *testing.T) {
	k := 5
	A := []int{3, 4, 3, 4, 3}
	B := []int{4, 5, 4, 5, 4}
	expect := 3
	runSample(t, k, A, B, expect)
}
