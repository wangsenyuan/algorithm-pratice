package main

import (
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 4}
	B := []int{1, 2}
	var expect int64 = 15
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0}
	B := []int{0}
	var expect int64 = 0
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 4, 3}
	B := []int{1, 2, 1}
	var expect int64 = 30
	runSample(t, A, B, expect)
}
