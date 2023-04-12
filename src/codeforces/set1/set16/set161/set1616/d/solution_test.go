package main

import "testing"

func runSample(t *testing.T, A []int, x int, expect int) {
	res := solve(A, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	x := 2
	expect := 4
	runSample(t, A, x, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 4, 2, 4, 2, 4, 2, 4, 2, 4}
	x := 3
	expect := 8
	runSample(t, A, x, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-10, -5, -10}
	x := 3
	expect := 2
	runSample(t, A, x, expect)
}

func TestSample4(t *testing.T) {
	A := []int{9, 9, -3}
	x := 5
	expect := 2
	runSample(t, A, x, expect)
}
