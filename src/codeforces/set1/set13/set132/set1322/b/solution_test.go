package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := bruteForce(a)
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	expect := bruteForce(a)
	runSample(t, a, expect)
}
