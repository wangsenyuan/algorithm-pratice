package main

import "testing"

func runSample(t *testing.T, a []int, m int, k int, expect int) {
	res := solve(a, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, k := 4, 2
	a := []int{2, 9, 2, 3, 8, 5}
	expect := 8
	runSample(t, a, m, k, expect)
}

func TestSample2(t *testing.T) {
	m, k := 4, 1
	a := []int{2, 13, 60, 4}
	expect := 4
	runSample(t, a, m, k, expect)
}

func TestSample3(t *testing.T) {
	m, k := 1, 3
	a := []int{1, 2, 2, 1}
	expect := 1
	runSample(t, a, m, k, expect)
}

func TestSample4(t *testing.T) {
	m, k := 2, 0
	a := []int{2, 2, 0}
	expect := 2
	runSample(t, a, m, k, expect)
}
