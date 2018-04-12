package main

import "testing"

func runSample(t *testing.T, n int, A, B string, expect int) {
	res := solve(n, []byte(A), []byte(B))

	if res != expect {
		t.Errorf("sample %d %s %s, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A, B := "ccadd", "bddcc"
	runSample(t, n, A, B, 3)
}
