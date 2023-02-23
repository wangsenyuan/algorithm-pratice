package main

import "testing"

func runSample(t *testing.T, k int, A, B []int, expect int) {
	res := solve(k, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 10
	A := []int{3, 4, 5}
	B := []int{2, 4, 2}
	expect := 2
	runSample(t, k, A, B, expect)
}

func TestSample2(t *testing.T) {
	k := 8
	A := []int{3, 4, 5}
	B := []int{2, 4, 2}
	expect := 1
	runSample(t, k, A, B, expect)
}

func TestSample3(t *testing.T) {
	k := 20
	A := []int{23, 54, 124, 54, 83}
	B := []int{2, 9, 5, 2, 10}
	expect := 0
	runSample(t, k, A, B, expect)
}

func TestSample4(t *testing.T) {
	k := 20
	A := []int{4, 7, 12, 34, 13}
	B := []int{30, 4, 3, 0, 9}
	expect := 2
	runSample(t, k, A, B, expect)
}
