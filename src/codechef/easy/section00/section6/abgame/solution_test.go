package main

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := solve(S)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "A.B", "A")
}

func TestSample2(t *testing.T) {
	runSample(t, "A..B", "A")
}

func TestSample3(t *testing.T) {
	runSample(t, "A..B.A..B", "B")
}

func TestSample4(t *testing.T) {
	runSample(t, "A...", "A")
}

func TestSample5(t *testing.T) {
	runSample(t, "B...", "B")
}
