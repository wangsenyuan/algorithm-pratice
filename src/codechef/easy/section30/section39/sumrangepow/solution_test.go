package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{2, 3, 2}
	expect := 9
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{2, 3, 2}
	expect := 15
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	A := []int{1, 1, 2, 2}
	expect := 38
	runSample(t, k, A, expect)
}

func TestSample4(t *testing.T) {
	k := 4
	A := []int{7, 3, 4, 1, 7, 3, 4}
	expect := 3068
	runSample(t, k, A, expect)
}

func TestSample5(t *testing.T) {
	k := 5
	A := []int{1, 2, 2, 6, 6, 2, 2, 1, 4, 4}
	expect := 13989
	runSample(t, k, A, expect)
}
