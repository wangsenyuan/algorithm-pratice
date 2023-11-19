package main

import "testing"

func runSample(t *testing.T, h int, x []int, expect int) {
	res := solve(h, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := 0
	x := []int{2, 7, 11, 13}
	expect := 0
	runSample(t, h, x, expect)
}

func TestSample2(t *testing.T) {
	h := 1
	x := []int{2, 7, 11, 13}
	expect := 3
	runSample(t, h, x, expect)
}

func TestSample3(t *testing.T) {
	h := 100
	x := []int{2, 7, 11, 13}
	expect := 107
	runSample(t, h, x, expect)
}

func TestSample4(t *testing.T) {
	h := 3
	x := []int{2, 5, 8, 11}
	expect := 8
	runSample(t, h, x, expect)
}
