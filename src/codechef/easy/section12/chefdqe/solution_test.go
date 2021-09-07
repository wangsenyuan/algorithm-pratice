package main

import "testing"

func runSample(t *testing.T, A []int, m int, expect int) {
	res := solve(m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 7
	A := []int{8, 6, 4, 6}
	expect := 1
	runSample(t, A, m, expect)
}

func TestSample2(t *testing.T) {
	m := 13
	A := []int{10, 2, 3, 8, 1, 10, 4}
	expect := 2
	runSample(t, A, m, expect)
}

func TestSample3(t *testing.T) {
	m := 11
	A := []int{3, 4, 8}
	expect := -1
	runSample(t, A, m, expect)
}
