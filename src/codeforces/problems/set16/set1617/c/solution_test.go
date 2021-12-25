package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 7}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 5, 4}
	expect := -1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{12345678, 87654321, 20211218, 23571113}
	expect := 4
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4, 18, 19, 5, 6, 7}
	expect := 2
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{2, 3}
	expect := 1
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{3, 4, 5}
	expect := 2
	runSample(t, A, expect)
}

func TestSample7(t *testing.T) {
	A := []int{5, 5, 5, 9, 14}
	expect := 4
	runSample(t, A, expect)
}
