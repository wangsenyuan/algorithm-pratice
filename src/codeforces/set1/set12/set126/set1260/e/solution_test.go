package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 9, 1, -1}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{11, -1, 13, 19, 24, 7, 17, 5}
	expect := 12
	runSample(t, a, expect)
}
