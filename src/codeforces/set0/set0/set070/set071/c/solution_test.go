package main

import "testing"

func runSample(t *testing.T, knights []int, expect bool) {
	res := solve(knights)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	knights := []int{1, 1, 1}
	expect := true
	runSample(t, knights, expect)
}

func TestSample2(t *testing.T) {
	knights := []int{1, 0, 1, 1, 1, 0}
	expect := true
	runSample(t, knights, expect)
}

func TestSample3(t *testing.T) {
	knights := []int{1, 0, 0, 1, 0, 1}
	expect := false
	runSample(t, knights, expect)
}
