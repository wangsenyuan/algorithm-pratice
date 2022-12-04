package main

import "testing"

func runSample(t *testing.T, P []int, expect int) {
	res := int(solve(P))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{4, 2, 1, 3}
	expect := 3
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{3, 1, 2}
	expect := 1
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{1, 6, 5, 2, 4, 3}
	expect := 3
	runSample(t, P, expect)
}

func TestSample4(t *testing.T) {
	P := []int{4, 2, 3, 1, 5}
	expect := 9
	runSample(t, P, expect)
}

func TestSample5(t *testing.T) {
	P := []int{3, 6, 4, 1, 2, 5}
	expect := 5
	runSample(t, P, expect)
}
