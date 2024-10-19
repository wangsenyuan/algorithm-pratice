package main

import "testing"

func runSample(t *testing.T, n int, k int, a []int, expect int) {
	res := solve(n, k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	n := 2
	k := 100
	a := []int{1410065408, 10000000000}
	expect := 1
	runSample(t, n, k, a, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	k := 12
	a := []int{2, 2}
	expect := 2
	runSample(t, n, k, a, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	k := 4
	a := []int{2, 6, 1, 2, 4}
	expect := 3
	runSample(t, n, k, a, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	k := 0
	a := []int{2, 1, 2}
	// 1, 1, 2, 3, 3
	expect := 1
	runSample(t, n, k, a, expect)
}

func TestSample5(t *testing.T) {
	n := 3
	k := 1
	a := []int{3, 2, 2}
	// 1, 1, 2, 3, 3
	expect := 2
	runSample(t, n, k, a, expect)
}
