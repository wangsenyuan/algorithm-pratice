package main

import "testing"

func runSample(t *testing.T, A []int, p int, expect int) {
	res := solve(A, p)

	if res != expect {
		t.Fatalf("Sample expect %d,but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 1}
	p := 4
	expect := 9
	runSample(t, A, p, expect)
}

func TestSample2(t *testing.T) {
	A := []int{20, 39, 5, 200}
	p := 7
	expect := 14
	runSample(t, A, p, expect)
}

func TestSample3(t *testing.T) {
	A := []int{48763, 1000000000}
	p := 200000
	expect := 448201910
	runSample(t, A, p, expect)
}
