package main

import "testing"

func runSample(t *testing.T, K int, S string, expect int) {
	res := solve(K, S)

	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", K, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 4
	S := "abcabcabcabc"
	expect := 7
	runSample(t, K, S, expect)
}

func TestSample2(t *testing.T) {
	K := 3
	S := "abcabcabcabc"
	expect := 12
	runSample(t, K, S, expect)
}
