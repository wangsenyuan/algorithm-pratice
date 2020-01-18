package main

import "testing"

func runSample(t *testing.T, s string, p int, expect string) {
	buf := []byte(s)
	solve(buf, p)

	if expect != string(buf) {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, string(buf))
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "bba", 3, "aab")
}
