package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 1, 4, 6, 7, 7, 5}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1}
	expect := 0
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 2, 2}
	expect := 3
	runSample(t, A, expect)
}
