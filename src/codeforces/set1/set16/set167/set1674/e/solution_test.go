package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{20, 10, 30, 10, 20}
	expect := 10
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 8, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 6, 6, 8, 5, 8}
	expect := 4
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 10}
	expect := 5
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{11, 1}
	expect := 6
	runSample(t, a, expect)
}
