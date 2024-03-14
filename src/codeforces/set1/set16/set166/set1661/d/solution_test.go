package main

import "testing"

func runSample(t *testing.T, b []int, k int, expect int) {
	res := solve(b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	b := []int{5, 4, 6}
	expect := 5
	runSample(t, b, k, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	b := []int{1, 2, 3, 2, 2, 3}
	expect := 3
	runSample(t, b, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	b := []int{1, 2, 4, 1, 2, 3}
	expect := 3
	runSample(t, b, k, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	b := []int{50, 17, 81, 25, 42, 39, 96}
	expect := 92
	runSample(t, b, k, expect)
}
