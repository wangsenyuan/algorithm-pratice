package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect bool) {
	res := solve(A, k)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	A := []int{2, 4, 5, 4, 2}
	expect := true
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	A := []int{4, 5, 5, 4, 6, 4}
	expect := true
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	A := []int{4, 5, 5, 4, 6, 4}
	expect := true
	runSample(t, k, A, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	A := []int{1, 2, 3, 3}
	expect := false
	runSample(t, k, A, expect)
}
