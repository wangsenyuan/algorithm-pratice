package main

import "testing"

func runSample(t *testing.T, n int, k int, B []int, expect bool) {
	res := solve(n, k, B)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	k := 3
	B := []int{1}
	expect := false
	runSample(t, n, k, B, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	k := 3
	B := []int{1, 5, 7}
	expect := true
	runSample(t, n, k, B, expect)
}
