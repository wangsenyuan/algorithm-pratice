package main

import "testing"

func runSample(t *testing.T, s string, K int, expect int) {
	res := solve([]byte(s), K)

	if res != expect {
		t.Errorf("Sample %s %d, expect %d, but got %d", s, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababc"
	K := 2
	expect := 6
	runSample(t, s, K, expect)
}
