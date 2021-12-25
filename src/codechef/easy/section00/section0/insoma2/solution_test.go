package main

import "testing"

func TestSample(t *testing.T) {
	var A, B, X byte
	A, B, X = 'A', 'B', 'X'
	n := 5
	matrix := [][]byte{
		{B, X, A, X, X},
		{X, A, X, X, X},
		{X, B, B, X, X},
		{X, A, X, A, X},
		{X, X, B, X, X},
	}
	s := []byte{A, B, A}

	res := solve(n, matrix, s)

	if res != 14 {
		t.Errorf("the sample should have 14 ways, but got %d", res)
	}
}
