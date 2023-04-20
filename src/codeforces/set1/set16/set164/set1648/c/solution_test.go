package main

import "testing"

func runSample(t *testing.T, S []int, R []int, expect int) {
	res := solve(S, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []int{1, 2, 2}
	R := []int{2, 1, 2, 1}
	expect := 2
	runSample(t, S, R, expect)
}

func TestSample2(t *testing.T) {
	S := []int{1, 2, 3, 4}
	R := []int{4, 3, 2, 1}
	expect := 23
	runSample(t, S, R, expect)
}

func TestSample3(t *testing.T) {
	S := []int{1, 1, 1, 2}
	R := []int{1, 1, 2}
	expect := 1
	runSample(t, S, R, expect)
}
