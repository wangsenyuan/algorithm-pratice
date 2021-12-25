package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	A := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expect := 1334961
	runSample(t, n, A, expect)
}
