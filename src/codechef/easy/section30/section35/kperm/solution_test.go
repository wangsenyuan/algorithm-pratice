package main

import "testing"

func runSample(t *testing.T, N int, K int, P []int, expect int) {
	res := solve(N, K, P)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

}

func TestSample1(t *testing.T) {
	n := 5
	k := 2
	P := []int{1, 0, 0, 0, 4}
	expect := 2
	runSample(t, n, k, P, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	k := 3
	P := []int{2, 0, 0, 0, 4}
	expect := 6
	runSample(t, n, k, P, expect)
}
