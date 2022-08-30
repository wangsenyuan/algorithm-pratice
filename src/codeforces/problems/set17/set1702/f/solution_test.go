package main

import "testing"

func runSample(t *testing.T, A, B []int, expect bool) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 4, 5, 24}
	B := []int{1, 4, 6, 11}
	expect := true
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 4, 17}
	B := []int{4, 5, 31}
	expect := false
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 7, 10, 13, 14}
	B := []int{2, 14, 14, 26, 42}
	expect := true
	runSample(t, A, B, expect)
}
func TestSample4(t *testing.T) {
	A := []int{2, 2, 4, 4, 4}
	B := []int{28, 46, 62, 71, 98}
	expect := true
	runSample(t, A, B, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 2, 10, 16, 64, 80}
	B := []int{20, 43, 60, 74, 85, 99}
	expect := true
	runSample(t, A, B, expect)
}
