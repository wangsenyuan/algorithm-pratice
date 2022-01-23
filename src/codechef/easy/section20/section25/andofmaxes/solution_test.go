package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{100, 10, 1, 1000, 10000}
	expect := 10000
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{7, 8, 9, 10, 11, 12}
	expect := 8
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	A := []int{26447356, 268435455, 56544987, 1000000000, 296823278}
	expect := 52087296
	runSample(t, k, A, expect)
}
