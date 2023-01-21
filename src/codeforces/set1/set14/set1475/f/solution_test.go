package main

import "testing"

func runSample(t *testing.T, n int, A []string, B []string, expect bool) {
	a := make([][]byte, n)
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i] = []byte(A[i])
		b[i] = []byte(B[i])
	}

	res := solve(n, a, b)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []string{
		"110",
		"001",
		"110",
	}
	B := []string{
		"000",
		"000",
		"000",
	}
	expect := true
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []string{
		"01",
		"11",
	}
	B := []string{
		"10",
		"10",
	}
	expect := false
	runSample(t, n, A, B, expect)
}
