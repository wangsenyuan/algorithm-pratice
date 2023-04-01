package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 0, 1}
	expect := 14
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 0, 5, 1}
	expect := 26
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 1, 1, 0, 1}
	expect := 48
	runSample(t, A, expect)
}
