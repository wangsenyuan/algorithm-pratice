package main

import "testing"

func runSample(t *testing.T, P, Q, R int, expect int) {
	res := int(solve(P, Q, R))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P, Q, R := 10, 12, 14
	expect := 4
	runSample(t, P, Q, R, expect)
}

func TestSample2(t *testing.T) {
	P, Q, R := 0, 5, 5
	expect := 1
	runSample(t, P, Q, R, expect)
}

func TestSample3(t *testing.T) {
	P, Q, R := 0, 0, 7
	expect := 0
	runSample(t, P, Q, R, expect)
}
