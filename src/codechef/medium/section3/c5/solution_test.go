package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(A, B)

	if len(res) != expect {
		t.Fatalf("Sample result %v, not correct", res)
	}

	if !isSubSeq(A, res) || !isSubSeq(B, res) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func isSubSeq(a, b []int) bool {
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			j++
		}
		i++
	}

	return j == len(b)
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 1, 4, 0}
	B := []int{10, 3, 4, 1, 0, 0}
	expect := 2
	runSample(t, A, B, expect)
}
