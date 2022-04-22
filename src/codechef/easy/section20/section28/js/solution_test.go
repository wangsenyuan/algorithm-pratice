package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 2, 3, 4}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 1, 3, 4}
	expect := 1
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, 3, 2, 4}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{1, 3, 4, 2, 4}
	expect := 3
	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	A := []int{1, 3, 4, 1, 1, 4}
	expect := 3
	runSample(t, n, A, expect)
}

func TestSample6(t *testing.T) {
	n := 7
	A := []int{1, 3, 4, 1, 1, 4, 5}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample7(t *testing.T) {
	n := 1
	A := []int{1}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample8(t *testing.T) {
	n := 2
	A := []int{1, 1}
	expect := 1
	runSample(t, n, A, expect)
}
