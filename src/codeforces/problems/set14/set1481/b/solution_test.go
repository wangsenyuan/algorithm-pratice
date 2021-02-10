package main

import "testing"

func runSample(t *testing.T, H []int, K int, expect int) {
	res := solve(H, K)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{4, 1, 2, 3}
	K := 3
	expect := 2
	runSample(t, H, K, expect)
}

func TestSample2(t *testing.T) {
	H := []int{1, 8}
	K := 7
	expect := 1
	runSample(t, H, K, expect)
}

func TestSample3(t *testing.T) {
	H := []int{4, 1, 2, 3}
	K := 5
	expect := -1
	runSample(t, H, K, expect)
}

func TestSample4(t *testing.T) {
	H := []int{5, 3, 1}
	K := 1
	expect := -1
	runSample(t, H, K, expect)
}
