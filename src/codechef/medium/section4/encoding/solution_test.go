package main

import "testing"

func runSample(t *testing.T, L, R string, expect int) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "9", "97", 4681)
}

func TestSample2(t *testing.T) {
	runSample(t, "8", "12", 49)
}

func TestSample3(t *testing.T) {
	runSample(t, "1", "8", 36)
}

func TestSample4(t *testing.T) {
	runSample(t, "1", "5432", 14630495)
}

func TestSample5(t *testing.T) {
	runSample(t, "1", "432", 92745)
}
