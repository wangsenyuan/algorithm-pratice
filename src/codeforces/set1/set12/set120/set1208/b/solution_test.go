package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 2, 2}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 4, 1, 4, 9}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{7, 6, 5, 4, 3, 2, 1, 7, 7, 7}
	expect := 3
	runSample(t, a, expect)
}