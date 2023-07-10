package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1, 1, 3, 5, 1}
	B := []int{3, 5}
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 2}
	B := []int{1, 2, 3}
	expect := -1
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 3, 1, 2, 3}

	expect := 2
	runSample(t, A, nil, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 3, 5, 6, 12, 9, 8, 10, 13, 15}
	B := []int{2, 4, 9}
	expect := 3
	runSample(t, A, B, expect)
}
