package main

import "testing"

func runSample(t *testing.T, n int, B []int, expect int) {
	res := solve(n, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	B := []int{1}
	expect := 1
	runSample(t, n, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	B := []int{1, 4, 5, 9, 10}
	expect := 3
	runSample(t, n, B, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	B := []int{3, 4}
	expect := 1
	runSample(t, n, B, expect)
}
