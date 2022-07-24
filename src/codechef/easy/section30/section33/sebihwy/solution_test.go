package main

import "testing"

func runSample(t *testing.T, S, SG, FG, D, T int, expect string) {
	res := solve(S, SG, FG, D, T)

	if res != expect {
		t.Errorf("Sample expect %s, bug got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S, SG, FG, D, T := 100, 180, 200, 20, 60
	expect := SEBI
	runSample(t, S, SG, FG, D, T, expect)
}

func TestSample2(t *testing.T) {
	S, SG, FG, D, T := 130, 131, 132, 1, 72
	expect := FATHER
	runSample(t, S, SG, FG, D, T, expect)
}
