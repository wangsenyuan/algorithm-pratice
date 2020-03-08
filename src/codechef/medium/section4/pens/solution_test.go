package main

import "testing"

func runSample(t *testing.T, n, k int, S string, P []string) {
	res := solve(n, k, S, P)
	t.Logf("%v", res)
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	S := "abcd"
	P := []string{
		"ab", "cd",
	}
	runSample(t, n, k, S, P)
}
