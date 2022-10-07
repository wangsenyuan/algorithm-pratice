package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := int(solve(A, B))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 26, -79, 72, 23}
	B := []int{66, 44}
	expect := 205
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{81}
	B := []int{-97}
	expect := 81
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{10, -5, 14, -20, 4}
	B := []int{-10, 5, -2}
	expect := 24
	runSample(t, A, B, expect)
}
