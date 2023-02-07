package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int64) {
	res := solve(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 6, 1, 2, 3}
	k := 2
	var expect int64 = 12
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 6, 1, 2, 3}
	k := 7
	var expect int64 = 37
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{999999}
	k := 2
	var expect int64 = 1000000
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	k := 70000
	var expect int64 = 5000349985
	runSample(t, A, k, expect)
}
