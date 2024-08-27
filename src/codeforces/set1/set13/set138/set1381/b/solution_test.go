package main

import "testing"

func runSample(t *testing.T, p []int, expect bool) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 3, 1, 4}
	expect := true
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{3, 1, 2, 4}
	expect := false
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{3, 2, 6, 1, 5, 7, 8, 4}
	expect := true
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{1, 2, 3, 4, 5, 6}
	expect := true
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{6, 1, 3, 7, 4, 5, 8, 2}
	expect := false
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{4, 3, 2, 5, 1, 11, 9, 12, 8, 6, 10, 7}
	expect := false
	runSample(t, p, expect)
}
