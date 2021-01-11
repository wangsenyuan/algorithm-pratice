package main

import "testing"

func runSample(t *testing.T, K int, H []int, expect int) {
	res := solve(K, H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 38
	H := []int{7, 8, 19, 7, 8, 7, 10, 20}
	expect := 7
	runSample(t, K, H, expect)
}

func TestSample2(t *testing.T) {
	K := 5
	H := []int{2, 10, 4, 9}
	expect := 2
	runSample(t, K, H, expect)
}
