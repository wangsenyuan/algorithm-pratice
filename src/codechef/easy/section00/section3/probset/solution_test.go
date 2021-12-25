package main

import "testing"

func runSample(t *testing.T, N, M int, plan []string, output []string, expect string) {
	res := solve(N, M, plan, output)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %s, but got %s", N, M, plan, output, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 4, 5
	plan := []string{"correct", "wrong", "correct", "wrong"}
	output := []string{"11111", "10101", "11111", "01111"}
	expect := "FINE"
	runSample(t, N, M, plan, output, expect)
}

func TestSample2(t *testing.T) {
	N, M := 4, 5
	plan := []string{"correct", "wrong", "correct", "wrong"}
	output := []string{"10110", "11111", "11111", "01000"}
	expect := "INVALID"
	runSample(t, N, M, plan, output, expect)
}

func TestSample3(t *testing.T) {
	N, M := 4, 5
	plan := []string{"correct", "wrong", "wrong", "wrong"}
	output := []string{"11111", "11111", "10100", "00000"}
	expect := "WEAK"
	runSample(t, N, M, plan, output, expect)
}

func TestSample4(t *testing.T) {
	N, M := 4, 5
	plan := []string{"wrong", "correct", "correct", "wrong"}
	output := []string{"10110", "01101", "11111", "11000"}
	expect := "INVALID"
	runSample(t, N, M, plan, output, expect)
}
