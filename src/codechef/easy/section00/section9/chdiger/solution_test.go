package main

import "testing"

func runSample(t *testing.T, s string, d byte, expect string) {
	res := solve([]byte(s), d)
	if res != expect {
		t.Errorf("Sample %s %c, expect %s, but got %s", s, d, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "35", '4', "34")
}

func TestSample2(t *testing.T) {
	runSample(t, "42", '4', "24")
}

func TestSample3(t *testing.T) {
	runSample(t, "24", '9', "24")
}
