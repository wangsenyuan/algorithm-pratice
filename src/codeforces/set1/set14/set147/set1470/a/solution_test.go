package main

import "testing"

func runSample(t *testing.T, K []int, C []int, expect int64) {
	res := solve(K, C)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := []int{2, 3, 4, 3, 2}
	C := []int{3, 5, 12, 20}
	var expect int64 = 30
	runSample(t, K, C, expect)
}

func TestSample2(t *testing.T) {
	K := []int{5, 4, 3, 2, 1}
	C := []int{10, 40, 90, 160, 250}
	var expect int64 = 190
	runSample(t, K, C, expect)
}

func TestSample3(t *testing.T) {
	K := []int{1}
	C := []int{1}
	var expect int64 = 1
	runSample(t, K, C, expect)
}

func TestSample4(t *testing.T) {
	K := []int{7, 3}
	C := []int{211047202, 262703497, 433933447, 476190629, 623690081, 861648772, 971407775}
	var expect int64 = 473750699
	runSample(t, K, C, expect)
}
