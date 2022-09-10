package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{3, 4}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	A := []int{1}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 1000000000
	A := []int{271832455, 357062697, 396505195, 580082912, 736850926}
	expect := 999999980
	runSample(t, n, A, expect)
}
