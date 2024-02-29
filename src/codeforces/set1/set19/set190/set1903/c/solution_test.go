package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, -3, 7, -6, 2, 5}
	expect := 32
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 9, -5, -3}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-3, -4, 2, -5, 1, 10, 17, 23}
	expect := 343
	runSample(t, a, expect)
}
