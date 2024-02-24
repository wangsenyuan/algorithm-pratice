package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 1, 3, 2}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1000, 1000}
	expect := 1
	runSample(t, a, expect)
}
func TestSample4(t *testing.T) {
	a := []int{1, 1, 1}
	expect := 3
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 4, 1, 6, 2, 8, 5, 4, 2, 10, 5, 10, 8, 7, 4, 3, 2, 6, 10}
	expect := 19
	runSample(t, a, expect)
}
