package main

import "testing"

func runSample(t *testing.T, n int, P1 []int, P2 []int, expect int) {
	res := solve(n, P1, P2)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	P1 := []int{-1}
	P2 := []int{-1}
	expect := 1
	runSample(t, n, P1, P2, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	P1 := []int{2, -1, 2}
	P2 := []int{2, -1, 1}
	expect := 2
	runSample(t, n, P1, P2, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	P1 := []int{2, 3, -1, 2, 2, 4, 4}
	P2 := []int{6, -1, 4, 5, 2, 4, 4}
	expect := 2
	runSample(t, n, P1, P2, expect)
}
