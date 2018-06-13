package main

import "testing"

func runSample(t *testing.T, n int, K int, A []int, expect int64) {
	res := solve(n, K, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 3, 2
	A := []int{1, 2, 2}
	runSample(t, n, K, A, 5)
}

func TestSample2(t *testing.T) {
	n, K := 4, 3
	A := []int{1, 2, 3, 4}
	runSample(t, n, K, A, 10)
}

func TestSample3(t *testing.T) {
	n, K := 2, 2
	A := []int{1, 2}
	runSample(t, n, K, A, 3)
}

func TestSample4(t *testing.T) {
	n, K := 11, 4
	A := []int{66, 152, 7, 89, 42, 28, 222, 69, 10, 54, 99}
	runSample(t, n, K, A, 704)
}
