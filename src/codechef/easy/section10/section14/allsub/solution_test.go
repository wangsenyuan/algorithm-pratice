package main

import "testing"

func runSample(t *testing.T, S, R string, expect string) {
	res := solve(S, R)

	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", S, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aa", "ababab", "aaabbb")
}

func TestSample2(t *testing.T) {
	runSample(t, "aaa", "ramialsadaka", "aaaaadiklmrs")
}

func TestSample3(t *testing.T) {
	runSample(t, "said", "sryhieni", "")
}

func TestSample4(t *testing.T) {
	runSample(t, "code", "codeisfun", "codefinsu")
}
