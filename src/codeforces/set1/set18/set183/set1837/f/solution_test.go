package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 10, 1, 1, 1}
	k := 4
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 20, 5, 15, 3}
	k := 3
	expect := 6
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 20, 3, 15, 5}
	k := 3
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{10, 8, 20, 14, 3, 8, 6, 4, 16, 11}
	k := 6
	expect := 21
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{9, 9, 2, 13, 15, 19, 4, 9, 13, 12}
	k := 5
	expect := 18
	runSample(t, a, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int{10, 9}
	k := 1
	expect := 9
	runSample(t, a, k, expect)
}
