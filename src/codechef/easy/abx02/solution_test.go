package main

import "testing"

func runSample(t *testing.T, P, Q, N, B1, B2, S1, S2 int, expect int) {
	res := solve(P, Q, N, B1, B2, S1, S2)
	if res != expect {
		t.Errorf("sample %d %d %d %d %d %d %d, should give %d, but got %d", P, Q, N, B1, B2, S1, S2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	P, Q, N, B1, B2, S1, S2 := 2, 2, 4, 2, 2, 1, 1
	expect := 2
	runSample(t, P, Q, N, B1, B2, S1, S2, expect)
}

func TestSample2(t *testing.T) {
	P, Q, N, B1, B2, S1, S2 := 10, 10, 10, 10, 10, 10, 10
	expect := 1024
	runSample(t, P, Q, N, B1, B2, S1, S2, expect)
}
