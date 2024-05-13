package main

import "testing"

func runSample(t *testing.T, n int, p []int, s []int, expect int) {
	res := solve(n, p, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	p := []int{1, 1, 1, 1}
	s := []int{1, -1, -1, -1, -1}
	expect := 1
	runSample(t, n, p, s, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	p := []int{1, 2, 3, 1}
	s := []int{1, -1, 2, -1, -1}
	expect := 2
	runSample(t, n, p, s, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	p := []int{1, 2}
	s := []int{2, -1, 1}
	expect := -1
	runSample(t, n, p, s, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	p := []int{1, 2, 2, 3}
	s := []int{1, -1, 2, 3, -1}
	expect := 3
	runSample(t, n, p, s, expect)
}
