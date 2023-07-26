package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := solve(A, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{0, 0, 0, 1}
	expect := -1
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	A := []int{3, 3, 3}
	expect := 1
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	A := []int{0, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	expect := 2
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	A := []int{1, 2, 3, 4, 5}
	expect := 2
	runSample(t, A, k, expect)
}

func TestSample5(t *testing.T) {
	k := 4
	A := []int{2, 2, 3, 5, 7, 11, 13, 13, 17}
	expect := 2
	runSample(t, A, k, expect)
}

func TestSample6(t *testing.T) {
	k := 7
	A := []int{0, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	expect := 1
	runSample(t, A, k, expect)
}
