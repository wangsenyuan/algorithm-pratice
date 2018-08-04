package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "115", "151")
}

func TestSample2(t *testing.T) {
	runSample(t, "1051", "1105")
}

func TestSample3(t *testing.T) {
	runSample(t, "6233", "6323")
}

func TestSample4(t *testing.T) {
	runSample(t, "511", "1015")
}
