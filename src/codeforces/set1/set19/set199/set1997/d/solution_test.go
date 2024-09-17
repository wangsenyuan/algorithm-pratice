package main

import "testing"

func runSample(t *testing.T, n int, a []int, p []int, expect int) {
	res := solve(n, a, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	a := []int{0, 1, 0, 2}
	p := []int{1, 1, 3}
	expect := 1
	runSample(t, n, a, p, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	a := []int{3, 0}
	p := []int{1}
	expect := 3
	runSample(t, n, a, p, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	a := []int{2, 5, 3, 9, 6}
	p := []int{3, 1, 5, 2}
	expect := 6
	runSample(t, n, a, p, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	a := []int{3, 1, 0}
	p := []int{1, 2}
	expect := 3
	runSample(t, n, a, p, expect)
}

func TestSample5(t *testing.T) {
	n := 3
	a := []int{3, 1, 1}
	p := []int{1, 2}
	expect := 4
	runSample(t, n, a, p, expect)
}

func TestSample6(t *testing.T) {
	n := 3
	a := []int{0, 3, 0}
	p := []int{1, 2}
	expect := 0
	runSample(t, n, a, p, expect)
}
