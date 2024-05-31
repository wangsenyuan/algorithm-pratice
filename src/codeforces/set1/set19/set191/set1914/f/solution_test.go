package main

import "testing"

func runSample(t *testing.T, n int, p []int, expect int) {
	res := solve(n, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	p := []int{1, 2, 1}
	expect := 1
	runSample(t, n, p, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	p := []int{1}
	expect := 0
	runSample(t, n, p, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	p := []int{5, 5, 5, 1}
	expect := 1
	runSample(t, n, p, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	p := []int{1, 2, 1, 1, 3, 3}
	expect := 3
	runSample(t, n, p, expect)
}

func TestSample5(t *testing.T) {
	n := 7
	p := []int{1, 1, 3, 2, 2, 4}
	expect := 3
	runSample(t, n, p, expect)
}

func TestSample6(t *testing.T) {
	n := 7
	p := []int{1, 2, 1, 1, 1, 3}
	expect := 3
	runSample(t, n, p, expect)
}
