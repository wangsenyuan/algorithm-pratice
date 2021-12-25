package main

import "testing"

func runSample(t *testing.T, n, K int, S []string, expect int) {
	res := solve(n, K, S)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, K, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 5, 2
	S := []string{
		"11001",
		"01010",
		"10101",
		"00110",
		"10101",
	}
	expect := 4
	runSample(t, n, K, S, expect)
}

func TestSample2(t *testing.T) {
	n, K := 3, 1
	S := []string{
		"110",
		"010",
		"011",
	}
	expect := -1
	runSample(t, n, K, S, expect)
}
