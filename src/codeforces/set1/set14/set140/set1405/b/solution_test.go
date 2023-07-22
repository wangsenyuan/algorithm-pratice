package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{-3, 5, -3, 1}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, -1}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-3, 2, -3, 4}
	expect := 4
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{-1, 1, 1, -1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{-5, 7, -6, -4, 17, -13, 4}
	expect := 8
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{-1000000000, -1000000000, -1000000000, 1000000000, 1000000000, 1000000000}
	expect := 3000000000
	runSample(t, A, expect)
}
