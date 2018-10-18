package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aaa", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "aca", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "aaaa", 5)
}

func TestSample4(t *testing.T) {
	runSample(t, "aabaa", 2)
}

func TestSample5(t *testing.T) {
	// aaabaaa(-1), aaabaaa(0), aaabaaa(1), aaababaa(3)
	runSample(t, "aabaaa", 4)
}

func TestSample6(t *testing.T) {
	// aabbbaa
	runSample(t, "aabbaa", 3)
}
