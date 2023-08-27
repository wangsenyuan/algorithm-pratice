package main

import "testing"

func runSample(t *testing.T, A, B, C []int, expect int) {
	res := solve(A, B, C)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7}
	B := []int{2, 3, 1, 7, 6, 5, 4}
	C := []int{2, 0, 1, 0, 0, 0, 0}
	expect := 4
	runSample(t, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1}
	B := []int{1}
	C := []int{0}
	expect := 1
	runSample(t, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 5, 2, 4, 6, 3}
	B := []int{6, 5, 3, 1, 4, 2}
	C := []int{6, 0, 0, 0, 0, 0}
	expect := 2
	runSample(t, A, B, C, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 6, 4, 7, 2, 3, 8, 5}
	B := []int{3, 2, 8, 1, 4, 5, 6, 7}
	C := []int{1, 0, 0, 7, 0, 3, 0, 5}
	expect := 2
	runSample(t, A, B, C, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 8, 6, 2, 4, 7, 9, 3, 10, 5}
	B := []int{1, 9, 2, 3, 4, 10, 8, 6, 7, 5}
	C := []int{1, 9, 2, 3, 4, 10, 8, 6, 7, 5}
	expect := 1
	runSample(t, A, B, C, expect)
}

func TestSample6(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7}
	B := []int{2, 3, 1, 7, 6, 5, 4}
	C := []int{0, 0, 0, 0, 0, 0, 0}
	expect := 8
	runSample(t, A, B, C, expect)
}
