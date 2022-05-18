package main

import "testing"

func runSample(t *testing.T, A []int, M int, expect int) {
	res := solve(A, M)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{12, 1}
	M := 18
	expect := 24
	runSample(t, A, M, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 5, 6}
	M := 5
	expect := 15
	runSample(t, A, M, expect)
}

func TestSample3(t *testing.T) {
	A := []int{79, 29, 80, 58, 80}
	M := 4
	expect := 162
	runSample(t, A, M, expect)
}

func TestSample4(t *testing.T) {
	A := []int{33, 46, 56}
	M := 20
	expect := 112
	runSample(t, A, M, expect)
}
