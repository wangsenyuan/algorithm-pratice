package main

import "testing"

func runSample(t *testing.T, K []int, H []int, expect int64) {
	res := solve(K, H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := []int{6}
	H := []int{4}
	var expect int64 = 10
	runSample(t, K, H, expect)
}

func TestSample2(t *testing.T) {
	K := []int{4, 5}
	H := []int{2, 2}
	var expect int64 = 6
	runSample(t, K, H, expect)
}

func TestSample3(t *testing.T) {
	K := []int{5, 7, 9}
	H := []int{2, 1, 2}
	var expect int64 = 7
	runSample(t, K, H, expect)
}

func TestSample4(t *testing.T) {
	K := []int{1, 2, 3}
	H := []int{1, 1, 1}
	var expect int64 = 3
	runSample(t, K, H, expect)
}

func TestSample5(t *testing.T) {
	K := []int{1, 2, 3}
	H := []int{1, 2, 2}
	var expect int64 = 6
	runSample(t, K, H, expect)
}
