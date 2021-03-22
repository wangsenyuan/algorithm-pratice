package main

import "testing"

func runSample(t *testing.T, K int, A []int, expect int64) {
	res := solve(len(A), K, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 0, 1, 0, 1}
	K := 3
	var expect int64 = 3
	runSample(t, K, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 0, 1}
	K := 4
	var expect int64 = 6
	runSample(t, K, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 0}
	K := 1
	var expect int64 = 2
	runSample(t, K, A, expect)
}
