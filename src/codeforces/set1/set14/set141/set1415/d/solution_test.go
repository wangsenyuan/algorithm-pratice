package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 5, 6, 8}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3}
	expect := -1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 4, 6, 20}
	expect := 2
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{11, 22, 71, 92}
	expect := 2
	runSample(t, A, expect)
}
