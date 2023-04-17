package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 5}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 3, 7}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1000000}
	expect := false
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{49, 3, 5, 34, 18, 9, 17, 20, 11, 19, 52,
		26, 50, 25, 8, 16, 13, 24, 40, 48, 12, 51, 6, 32, 41, 10, 1, 2, 27, 4}
	expect := true
	runSample(t, A, expect)
}
