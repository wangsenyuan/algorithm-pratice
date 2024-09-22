package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 3, 0, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 2, 3, 4, 1, 1, 0, 2, 2}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 2, 1, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 0, 0, 4, 2, 2, 1}
	expect := true
	runSample(t, a, expect)
}
