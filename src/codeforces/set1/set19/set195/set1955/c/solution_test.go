package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 4, 3}
	k := 5
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 4, 3}
	k := 6
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 7, 1, 8, 2}
	k := 20
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 2}
	k := 2
	expect := 0
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 5}
	k := 15
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int{5, 2}
	k := 7
	expect := 2
	runSample(t, a, k, expect)
}
