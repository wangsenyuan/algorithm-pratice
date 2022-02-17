package main

import "testing"

func runSample(t *testing.T, s string, p int, expect string) {
	res := solve(s, p)

	if expect != res {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "bba", 3, "aab")
}

func TestSample2(t *testing.T) {
	runSample(t, "b", 2, "a")
}

func TestSample3(t *testing.T) {
	runSample(t, "bb", 2, "ab")
}

func TestSample4(t *testing.T) {
	runSample(t, "abbbaaaaabaaaaabaaaab", 11, "aaaaaaaaaaaaaaaaaaaab")
}
