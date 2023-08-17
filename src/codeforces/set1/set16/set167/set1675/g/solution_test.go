package main

import "testing"

func runSample(t *testing.T, n int, m int, A []int, expect int) {
	res := solve(n, m, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 19
	A := []int{5, 3, 2, 3, 3, 3}
	expect := 2
	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 6
	A := []int{3, 2, 1}
	expect := 0
	runSample(t, n, m, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 6
	A := []int{2, 1, 3}
	expect := 1
	runSample(t, n, m, A, expect)
}

func TestSample4(t *testing.T) {
	n, m := 6, 19
	A := []int{3, 4, 3, 3, 5, 1}
	expect := 3
	runSample(t, n, m, A, expect)
}

func TestSample5(t *testing.T) {
	n, m := 10, 1
	A := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	expect := 9
	runSample(t, n, m, A, expect)
}
