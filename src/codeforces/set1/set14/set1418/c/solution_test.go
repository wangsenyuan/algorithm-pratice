package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	A := []int{1, 0, 1, 1, 0, 1, 1, 1}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 1, 1, 1, 0}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	A := []int{1, 1, 1, 1, 0, 0, 1}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	A := []int{1, 1, 1, 1, 1, 1}
	expect := 2
	runSample(t, n, A, expect)
}
func TestSample5(t *testing.T) {
	n := 1
	A := []int{1}
	expect := 1
	runSample(t, n, A, expect)
}
