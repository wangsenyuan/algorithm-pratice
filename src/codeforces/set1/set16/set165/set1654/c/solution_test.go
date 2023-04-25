package main

import "testing"

func runSample(t *testing.T, C []int, expect bool) {
	res := solve(C)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{1, 2, 3}
	expect := true
	runSample(t, C, expect)
}

func TestSample2(t *testing.T) {
	C := []int{1, 1, 1, 1, 1, 1}
	expect := true
	runSample(t, C, expect)
}

func TestSample3(t *testing.T) {
	C := []int{2, 16, 1, 8, 64, 1, 4, 32}
	expect := true
	runSample(t, C, expect)
}

func TestSample4(t *testing.T) {
	C := []int{1, 2, 4, 7, 1, 1, 1, 1, 7, 2}
	expect := true
	runSample(t, C, expect)
}
