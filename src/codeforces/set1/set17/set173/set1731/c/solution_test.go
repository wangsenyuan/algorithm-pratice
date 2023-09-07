package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 2}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 2, 1, 5, 3}
	expect := 11
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 4, 4, 4}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 7, 3, 7, 1, 7, 3}
	expect := 20
	runSample(t, a, expect)
}
