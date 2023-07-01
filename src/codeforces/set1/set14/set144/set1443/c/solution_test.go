package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 7, 4, 5}
	B := []int{2, 1, 2, 4}
	expect := 5
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4}
	B := []int{3, 3, 3, 3}
	expect := 3
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2}
	B := []int{10, 10}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10, 10}
	B := []int{1, 2}
	expect := 3
	runSample(t, A, B, expect)
}
