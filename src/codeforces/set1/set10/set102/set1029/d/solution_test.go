package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 11
	a := []int{45, 1, 10, 12, 11, 7}
	expect := 7
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := []int{2, 78, 4, 10}
	expect := 12
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	a := []int{3, 7, 19, 3, 3}
	expect := 0
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	k := 1000000000
	a := []int{1000000000, 1000000000}
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	k := 994305000
	a := []int{922337204, 1000000000}
	expect := 1
	runSample(t, a, k, expect)
}
