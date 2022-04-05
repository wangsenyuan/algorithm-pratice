package main

import "testing"

func runSample(t *testing.T, A []int, B []int, C int, expect bool) {
	res := solve(A, B, C)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1}
	B := []int{1, 2}
	C := 1
	expect := true
	runSample(t, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1}
	B := []int{1, 2}
	C := 2
	expect := true
	runSample(t, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 9}
	B := []int{1, 2, 4}
	C := 5
	expect := false
	runSample(t, A, B, C, expect)
}
