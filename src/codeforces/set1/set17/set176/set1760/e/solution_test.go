package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 0, 1, 0}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 0, 0, 1, 0}
	expect := 7
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 0}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 0, 1, 1, 0, 0, 0, 1}
	expect := 13
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 1, 1}
	expect := 2
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1}
	expect := 21
	runSample(t, a, expect)
}
