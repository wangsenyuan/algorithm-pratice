package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 6
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 4, 3}
	expect := 7
	runSample(t, a, expect)
}
