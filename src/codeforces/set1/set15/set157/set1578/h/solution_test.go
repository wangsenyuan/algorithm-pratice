package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "()", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "()->()", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "()->()->()", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, "(()->())->()", 2)
}

func TestSample5(t *testing.T) {
	runSample(t, "()->(((()->())->()->())->())", 3)
}
