package main

import "testing"

func runSample(t *testing.T, m int, A []int, expect int) {
	res := solve(A, m)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	m := 6
	expect := 16
	runSample(t, m, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	m := 5
	expect := 17
	runSample(t, m, A, expect)
}
