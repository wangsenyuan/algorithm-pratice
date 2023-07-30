package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 100, 10, 30, 5}
	B := []int{101, 104, 105, 110, 130, 200}
	expect := 0
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 2, 2, 3, 3}
	B := []int{13, 4, 11, 12, 11, 13, 12}
	expect := 7
	runSample(t, A, B, expect)
}
