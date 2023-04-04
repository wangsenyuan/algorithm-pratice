package main

import "testing"

func runSample(t *testing.T, A []int, x int, expect int) {
	res := solve(A, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 16, 4, 4}
	x := 4
	expect := 0
	runSample(t, A, x, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 2, 2, 4, 7}
	x := 2
	expect := 2
	runSample(t, A, x, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 2, 3, 5, 15}
	x := 3
	expect := 3
	runSample(t, A, x, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10, 10, 10, 20, 1, 100, 200, 2000, 3}
	x := 10
	expect := 3
	runSample(t, A, x, expect)
}
