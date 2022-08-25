package main

import "testing"

func runSample(t *testing.T, n int, A, B, G int, expect string) {
	res := solve(int64(n), int64(A), int64(B), int64(G))

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := 11
	B := 111
	G := 92
	expect := "084"
	runSample(t, n, A, B, G, expect)
}
