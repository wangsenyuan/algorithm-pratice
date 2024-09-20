package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k []int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 20, 30, 5}
	b := []int{9, 33, 115, 3}
	k := []int{2, 1, 1, 2}
	expect := 32
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{40, 1000, 300}
	b := []int{1, 1100, 2}
	k := []int{300, 2, 1}
	expect := 1337
	runSample(t, a, b, k, expect)
}
