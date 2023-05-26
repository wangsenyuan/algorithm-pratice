package main

import "testing"

func runSample(t *testing.T, B []int, x int, expect bool) {
	res := solve(B, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{3, 2, 1}
	x := 3
	expect := false
	runSample(t, B, x, expect)
}

func TestSample2(t *testing.T) {
	B := []int{1, 2, 3, 4}
	x := 3
	expect := true
	runSample(t, B, x, expect)
}

func TestSample3(t *testing.T) {
	B := []int{5, 1, 2, 3, 4}
	x := 2
	expect := true
	runSample(t, B, x, expect)
}
