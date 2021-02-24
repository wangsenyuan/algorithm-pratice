package main

import "testing"

func runSample(t *testing.T, n int, m int, A []int, B []int, expect int) {
	res := solve(n, m, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 7
	A := []int{5, 3, 2, 1, 4}
	B := []int{2, 1, 1, 2, 1}
	expect := 2
	runSample(t, n, m, A, B, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 3
	A := []int{2}
	B := []int{1}
	expect := -1
	runSample(t, n, m, A, B, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 10
	A := []int{2, 3, 2, 3, 2}
	B := []int{1, 2, 1, 2, 1}
	expect := 6
	runSample(t, n, m, A, B, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 10
	A := []int{5, 1, 3, 4}
	B := []int{1, 2, 1, 2}
	expect := 4
	runSample(t, n, m, A, B, expect)
}

func TestSample5(t *testing.T) {
	n, m := 4, 5
	A := []int{3, 2, 1, 2}
	B := []int{2, 1, 2, 1}
	expect := 3
	runSample(t, n, m, A, B, expect)
}
