package main

import (
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 1, 2, 1}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 11
	A := []int{2, 2, 2, 10, 2, 10, 1, 1, 5, 3, 5}
	expect := 3
	runSample(t, n, A, expect)
}
