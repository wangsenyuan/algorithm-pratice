package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4, 4, 2}
	expect := 8
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 3, 5}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{100003, 100004, 100005, 100006}
	expect := 10000800015
	runSample(t, a, expect)
}
