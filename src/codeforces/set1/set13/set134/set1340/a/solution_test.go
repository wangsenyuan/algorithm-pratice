package main

import "testing"

func runSample(t *testing.T, p []int, expect bool) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 3, 4, 5, 1}
	expect := true
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1}
	expect := true
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{1, 3, 2}
	expect := false
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{4, 2, 3, 1}
	expect := true
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{1, 5, 2, 4, 3}
	expect := false
	runSample(t, p, expect)
}
