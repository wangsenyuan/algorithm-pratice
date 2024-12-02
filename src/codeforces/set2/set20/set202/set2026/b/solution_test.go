package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 9}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 5, 8, 10, 13}
	expect := 3
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 2, 3, 5}
	expect := 2
	runSample(t, a, expect)
}
