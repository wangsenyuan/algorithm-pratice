package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 4, 4}
	B := []int{5, 5, 5}
	expect := 300
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 5, 5}
	B := []int{10, 8, 6}
	expect := 408
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 6, 7}
	B := []int{5, 6, 7}
	expect := 240
	runSample(t, A, B, expect)
}
