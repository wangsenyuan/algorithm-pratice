package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := []int{2, 1, 3}
	expect := 1
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := []int{2, 1}
	expect := 0
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := []int{1, 2, 3}
	expect := 1
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := []int{2, 1, 4, 3}
	expect := 2
	runSample(t, n, expect)
}
