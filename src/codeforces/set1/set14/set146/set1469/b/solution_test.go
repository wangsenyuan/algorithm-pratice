package main

import "testing"

func runSample(t *testing.T, R []int, B []int, expect int) {
	res := solve(R, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	R := []int{6, -5, 7, -3}
	B := []int{2, 3, -4}
	expect := 13
	runSample(t, R, B, expect)
}

func TestSample2(t *testing.T) {
	R := []int{1, 1}
	B := []int{10, -3, 2, 2}
	expect := 13
	runSample(t, R, B, expect)
}

func TestSample3(t *testing.T) {
	R := []int{2, 1, -3, 1, 0, 2, 0, 0}
	B := []int{3, -3, -4, -3, 4, 3, 1}
	expect := 6
	runSample(t, R, B, expect)
}
