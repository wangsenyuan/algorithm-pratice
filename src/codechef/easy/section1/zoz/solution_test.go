package main

import "testing"

func runSample(t *testing.T, n int, A []int, K int, expect int) {
	res := solve(n, A, K)
	if res != expect {
		t.Errorf("sample %d %v %d, expect %d, but got %d", n, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 4, 4
	A := []int{2, 1, 6, 7}
	expect := 1
	runSample(t, n, A, K, expect)
}

func TestSample2(t *testing.T) {
	n, K := 4, 2
	A := []int{2, 1, 5, 4}
	expect := 0
	runSample(t, n, A, K, expect)
}
