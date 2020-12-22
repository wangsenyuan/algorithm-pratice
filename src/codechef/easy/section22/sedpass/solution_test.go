package main

import "testing"

func runSample(t *testing.T, S string, expect int64) {
	res := solve(S)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "aa?"
	var expect int64 = 2
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "a???"
	var expect int64 = 6
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "????"
	var expect int64 = 4
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "asfhaslskfak"
	var expect int64 = 2
	runSample(t, S, expect)
}

func TestSample5(t *testing.T) {
	S := "af??avvnfed?fav?faf????"
	var expect int64 = 27
	runSample(t, S, expect)
}
