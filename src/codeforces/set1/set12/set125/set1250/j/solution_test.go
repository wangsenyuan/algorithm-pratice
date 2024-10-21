package main

import "testing"

func runSample(t *testing.T, k int, c []int, expect int) {
	res := solve(k, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 4
	c := []int{7, 1, 13}
	expect := 16
	runSample(t, k, c, expect)
}

func TestSample2(t *testing.T) {
	k := 100
	c := []int{100}
	expect := 100
	runSample(t, k, c, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	c := []int{1000000000000, 1000000000000}
	expect := 2000000000000
	runSample(t, k, c, expect)
}
