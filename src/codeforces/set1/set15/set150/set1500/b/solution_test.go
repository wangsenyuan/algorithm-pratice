package main

import "testing"

func runSample(t *testing.T, A, B []int, K int64, expect int64) {
	res := solve(A, B, K)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 3, 1}
	B := []int{2, 1}
	var K int64 = 4
	var expect int64 = 5
	runSample(t, A, B, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 2}
	B := []int{1, 6, 4, 3, 5, 7, 2, 8}
	var K int64 = 41
	var expect int64 = 47
	runSample(t, A, B, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2}
	B := []int{1}
	var K int64 = 31
	var expect int64 = 62
	runSample(t, A, B, K, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1}
	B := []int{2, 1}
	var K int64 = 1
	var expect int64 = 1
	runSample(t, A, B, K, expect)
}
