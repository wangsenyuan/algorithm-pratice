package main

import "testing"

func runSample(t *testing.T, n, p int, K []int, expect int) {
	res := solve(n, p, K)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, p, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, p := 5, 2
	K := []int{2, 3, 4, 4, 3}
	expect := 4
	runSample(t, n, p, K, expect)
}

func TestSample2(t *testing.T) {
	n, p := 3, 1
	K := []int{2, 10, 1000}
	expect := 1
	runSample(t, n, p, K, expect)
}

func TestSample3(t *testing.T) {
	n, p := 4, 5
	K := []int{0, 1, 1, 100}
	expect := 146981438
	runSample(t, n, p, K, expect)
}

func TestSample4(t *testing.T) {
	n, p := 1, 8
	K := []int{89}
	expect := 747093407
	runSample(t, n, p, K, expect)
}
