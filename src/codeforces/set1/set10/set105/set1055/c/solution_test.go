package main

import "testing"

func runSample(t *testing.T, alice []int, bob []int, expect int) {
	res := solve(alice, bob)

	if res != expect {
		t.Fatalf("sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 2, 5}
	b := []int{1, 3, 5}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 3}
	b := []int{2, 3, 6}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 3, 20}
	b := []int{6, 8, 10}
	expect := 0
	runSample(t, a, b, expect)
}
