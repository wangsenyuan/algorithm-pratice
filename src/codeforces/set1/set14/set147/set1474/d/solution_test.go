package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 2, 1, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2100, 1900, 1600, 3000, 1600}
	expect := true
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{2443, 2445}
	expect := false
	runSample(t, A, expect)
}
