package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 6, 15, 30}
	expect := 30
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 6, 30, 15}
	expect := 31
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 4, 6}
	expect := 10
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{30, 30, 15, 10, 6}
	expect := 31
	runSample(t, a, expect)
}
