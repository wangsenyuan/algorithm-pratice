package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 5, 4, 6}
	expect := 5
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	expect := 45
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{885, 8, 1, 128, 83, 32, 256, 206, 639, 16, 4, 128, 689, 32, 8, 64, 885, 969, 1}
	expect := 37
	runSample(t, a, expect)
}
