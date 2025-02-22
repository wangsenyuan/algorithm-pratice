package main

import "testing"

func runSample(t *testing.T, n int, A []string, expect bool) {
	res := solve(n, A)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	a, b, c := res[0], res[1], res[2]
	a--
	b--
	c--
	if A[a][b] != '1' || A[b][c] != '1' || A[c][a] != '1' {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []string{
		"00100",
		"10000",
		"01001",
		"11101",
		"11000",
	}
	expect := true
	runSample(t, n, A, expect)
}
