package main

import "testing"

func runSample(t *testing.T, m int, P []int, A []int, expect int) {
	res := solve(m, P, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	P := []int{2, 1, 4, 3}
	A := []int{0, 2, 0, 2}
	expect := 9
	runSample(t, m, P, A, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	P := []int{3, 1, 2}
	A := []int{0, 0, 0}
	expect := 8
	runSample(t, m, P, A, expect)
}

func TestSample3(t *testing.T) {
	m := 54
	P := []int{8, 1, 2, 4, 3, 6, 7, 5}
	A := []int{0, 0, 0, 0, 0, 0, 0, 0}
	expect := 459165024
	runSample(t, m, P, A, expect)
}
