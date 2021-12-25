package main

import "testing"

func runSample(t *testing.T, n int, m int, P []int, expect int) {
	res := solve(n, m, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 1
	P := []int{5}
	expect := 3
	runSample(t, n, m, P, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1000000000, 1
	P := []int{998244353}
	expect := 1755647
	runSample(t, n, m, P, expect)
}
