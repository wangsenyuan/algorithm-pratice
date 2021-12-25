package main

import "testing"

func runSample(t *testing.T, K int, A []int, expect int) {
	res := solve(int64(K), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 4
	A := []int{1, 0, 2}
	expect := 1
	runSample(t, K, A, expect)
}

func TestSample2(t *testing.T) {
	K := 2
	A := []int{1, 3, 2}
	expect := 0
	runSample(t, K, A, expect)
}

func TestSample3(t *testing.T) {
	K := 6
	A := []int{1, 0, 2}
	expect := 3
	runSample(t, K, A, expect)
}
