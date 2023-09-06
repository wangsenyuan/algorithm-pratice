package main

import "testing"

func runSample(t *testing.T, A []int, m int, expect int) {
	res := solve(A, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 0, 1}
	m := 2
	expect := 1
	runSample(t, A, m, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1}
	m := 2
	expect := 0
	runSample(t, A, m, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 1, 0}
	m := 2
	expect := 2
	runSample(t, A, m, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 0, 1, 2, 0, 1, 0}
	m := 3
	expect := 2
	runSample(t, A, m, expect)
}

func TestSample5(t *testing.T) {
	A := []int{0, 0, 1, 0, 0, 2, 1, 0}
	m := 3
	expect := 1
	runSample(t, A, m, expect)
}
