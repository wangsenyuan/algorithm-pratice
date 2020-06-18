package main

import "testing"

func runSample(t *testing.T, N, K int, A []int, expect int64) {
	res := solve(N, K, A)

	if res != expect {
		t.Errorf("Sample #{N} %d %v, expect %d, but got %d", K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 4, 2
	A := []int{1, 2, 3, 4}
	expect := int64(2)
	runSample(t, N, K, A, expect)
}
