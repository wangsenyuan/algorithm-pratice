package main

import (
	"testing"
)

func runSample(t *testing.T, A []int64, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{2, 4, 5}
	expect := int64(1)
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{1, 1, 1}
	expect := int64(-1)
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int64{1, 1, 2}
	expect := int64(-1)
	runSample(t, A, expect)
}
