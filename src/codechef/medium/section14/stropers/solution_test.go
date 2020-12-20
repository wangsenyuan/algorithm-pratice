package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []string{
		"00000",
		"10001",
		"10101",
		"01111",
		"11001",
		"01101",
		"10110",
		"10010",
		"10111",
		"11001",
	}
	expect := []int{
		5,
		11,
		8,
		9,
		12,
		10,
		10,
		11,
		11,
		12,
	}
	for i := 0; i < len(S); i++ {
		runSample(t, S[i], expect[i])
	}
}
