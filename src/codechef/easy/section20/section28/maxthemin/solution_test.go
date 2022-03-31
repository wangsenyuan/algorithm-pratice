package main

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := solve(A, K)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 4
	A := []int{5, 4, 3, 2, 1}
	expect := 5
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	K := 1
	A := []int{8, 2, 3, 9, 6, 10}
	expect := 3
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	K := 2
	A := []int{4, 3, 4, 3, 5}
	expect := 4
	runSample(t, A, K, expect)
}

func TestSample4(t *testing.T) {
	K := 1000
	A := []int{4, 3, 4, 3, 5}
	expect := 5
	runSample(t, A, K, expect)
}