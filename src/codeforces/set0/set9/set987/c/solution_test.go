package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 4, 5, 4, 10}
	B := []int{40, 30, 20, 10, 40}
	expect := 90
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{100, 101, 100}
	B := []int{2, 4, 5}
	expect := -1
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	B := []int{10, 13, 11, 14, 15, 12, 13, 13, 18, 13}
	expect := 33
	runSample(t, A, B, expect)
}
