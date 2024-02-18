package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 3, 2, 1}
	expect := 7
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, 5, 4, 5, 4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 7, 7, 2, 3, 4, 3, 2, 1, 100}
	expect := 28
	runSample(t, a, expect)
}
