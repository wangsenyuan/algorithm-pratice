package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 6}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-2, 4, -2, 4}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-10000000, -9999388, -9997084, -9988380, 10000000}
	expect := 13924
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 8, 8, 9, 9}
	expect := 2
	runSample(t, A, expect)
}
