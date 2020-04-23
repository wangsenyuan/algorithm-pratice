package main

import "testing"

func runSample(t *testing.T, B []int, expect int) {
	res := solve(len(B), B)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{2, 3}
	expect := 2
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{2, 6, 7, 7}
	expect := 64
	runSample(t, B, expect)
}
