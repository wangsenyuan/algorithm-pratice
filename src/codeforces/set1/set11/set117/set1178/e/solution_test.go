package main

import "testing"

func runSample(t *testing.T, s string) {
	res := solve(s)
	if len(res) < len(s)/2 {
		t.Fatalf("Sample result %s, not correct", res)
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		if res[i] != res[j] {
			t.Fatalf("Sample result %s, not a palindrome", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "cacbac"
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := "abc"
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := "cbacacacbcbababacbcb"
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := "abababacacacbcbcbc"
	runSample(t, s)
}
