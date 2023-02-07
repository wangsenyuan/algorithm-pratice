package main

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := int(solve(A, int64(K)))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{20}
	K := 10
	expect := 10
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{6, 9}
	K := 69
	expect := 0
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 1, 3, 1, 2, 1}
	K := 8
	expect := 2
	runSample(t, A, K, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 1, 2, 6, 1, 6, 8, 10}
	K := 1
	expect := 7
	runSample(t, A, K, expect)
}

/**
2 1
1 1*/
func TestSample5(t *testing.T) {
	A := []int{1, 1}
	K := 1
	expect := 1
	runSample(t, A, K, expect)
}
