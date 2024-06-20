package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{14, 25, 62, 74, 86, 95, 12}
	b := []int{51, 62, 71, 72, 92, 20, 84}
	expect := 30
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 0, 0, 100, 100, 100}
	b := []int{100, 100, 100, 0, 0, 0}
	expect := 100
	runSample(t, a, b, expect)
}
