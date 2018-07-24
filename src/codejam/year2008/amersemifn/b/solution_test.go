package main

import "testing"

func runSample(t *testing.T, n int, s []int, expect int) {
	res := solve(n, s)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	s := []int{1, 2, 3, 4, 5, 6, 7}
	expect := -1
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	s := []int{1, 10, 11, 200}
	expect := 201
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	s := []int{1000, 1520, 7520, 7521}
	expect := 3514
	runSample(t, n, s, expect)
}
