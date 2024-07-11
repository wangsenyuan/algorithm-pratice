package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 4, 2}
	k := 3
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{16, 192, 152, 78, 224, 202, 186, 52, 118, 19, 13, 38, 199, 196, 35, 295, 100, 64, 205, 37, 166, 124, 169, 214, 66, 243, 134, 192, 253, 270, 92}
	k := 3767
	// 1, 2, 2, 2, 2, 2  3, 4
	expect := 7546
	runSample(t, a, k, expect)
}
