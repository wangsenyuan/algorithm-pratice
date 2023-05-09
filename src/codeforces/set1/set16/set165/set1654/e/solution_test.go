package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 7, 8, 6, 9, 5, 4, 1}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{19, 2, 15, 8, 9, 14, 17, 13, 4, 14, 4, 11, 15, 7}
	expect := 10
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{100000, 1, 60000, 2, 20000, 4, 8, 16, 32, 64}
	expect := 7
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10000, 20000, 10000, 1}
	expect := 2
	runSample(t, A, expect)
}
