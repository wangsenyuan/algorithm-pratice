package main

import "testing"

func runSample(t *testing.T, A []int, L, R int, expect int) {
	res := solve(A, int64(L), int64(R))

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L, R := 10, 52
	A := []int{2, 5}
	expect := 3
	runSample(t, A, L, R, expect)
}

func TestSample2(t *testing.T) {
	L, R := 58, 100
	A := []int{4, 2, 3}
	expect := 0
	runSample(t, A, L, R, expect)
}

func TestSample3(t *testing.T) {
	L, R := 100, 1000
	A := []int{1, 10, 100, 1000}
	expect := 2
	runSample(t, A, L, R, expect)
}

func TestSample4(t *testing.T) {
	L, R := 28, 102
	A := []int{3, 2, 1, 9, 10}
	expect := 11
	runSample(t, A, L, R, expect)
}
