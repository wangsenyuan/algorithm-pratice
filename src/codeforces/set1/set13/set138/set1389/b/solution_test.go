package main

import "testing"

func runSample(t *testing.T, a []int, k int, z int, expect int) {
	res := solve(a, k, z)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 5, 4, 3, 2}
	k := 4
	z := 0
	expect := 15
	runSample(t, a, k, z, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 5, 4, 3, 2}
	k := 4
	z := 1
	expect := 19
	runSample(t, a, k, z, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 20, 30, 40, 50}
	k := 4
	z := 4
	expect := 150
	runSample(t, a, k, z, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 6, 8, 2, 9, 9, 7, 4, 10, 9}
	k := 7
	z := 3
	expect := 56
	runSample(t, a, k, z, expect)
}
