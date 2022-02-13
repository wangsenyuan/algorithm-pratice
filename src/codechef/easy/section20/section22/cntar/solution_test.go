package main

import "testing"

func runSample(t *testing.T, m int, A []int, expect int) {
	res := solve(m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	A := []int{2, 1}
	expect := 6
	runSample(t, m, A, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	A := []int{2, 1, 1}
	expect := 2
	runSample(t, m, A, expect)
}
