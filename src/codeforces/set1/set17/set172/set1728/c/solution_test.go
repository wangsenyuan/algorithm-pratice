package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 9, 3}
	B := []int{1, 100, 9}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{75019, 709259, 5, 611271314, 9024533, 81871864, 9, 3, 6, 4865}
	B := []int{9503, 2, 371245467, 6, 7, 37376159, 8, 364036498, 52295554, 169}
	expect := 18
	runSample(t, A, B, expect)
}
