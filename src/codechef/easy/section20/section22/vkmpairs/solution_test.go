package main

import "testing"

func runSample(t *testing.T, A []int, B []int, P int, expect int) {
	res := int(solve(A, B, P))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{7, 14}
	B := []int{2, 13}
	P := 17
	expect := 1
	runSample(t, A, B, P, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 91, 99, 200}
	B := []int{81, 182, 149}
	P := 101
	expect := 3
	runSample(t, A, B, P, expect)
}
