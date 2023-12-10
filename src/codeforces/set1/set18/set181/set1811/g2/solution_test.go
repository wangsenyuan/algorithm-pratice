package main

import "testing"

func runSample(t *testing.T, p []int, k int, expect int) {
	res := solve(p, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	p := []int{1, 2, 3, 4, 5}
	expect := 1
	runSample(t, p, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	p := []int{1, 3, 1, 3, 3, 1, 3}
	expect := 4
	runSample(t, p, k, expect)
}

func TestSample3(t *testing.T) {
	k := 4
	p := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	expect := 165
	runSample(t, p, k, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	p := []int{1, 1, 2, 2, 2}
	expect := 3
	runSample(t, p, k, expect)
}

func TestSample5(t *testing.T) {
	k := 1
	p := []int{1, 2, 3, 4, 5}
	expect := 1
	runSample(t, p, k, expect)
}
