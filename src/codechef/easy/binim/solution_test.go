package main

import "testing"

func runSample(t *testing.T, n int, stones []string, firstPlayer string, expect bool) {
	ss := make([][]byte, len(stones))
	for i := 0; i < len(stones); i++ {
		ss[i] = []byte(stones[i])
	}

	res := solve(n, ss, []byte(firstPlayer))

	if res != expect {
		t.Errorf("Sample %d %v %s, expect %t, but got %t", n, stones, firstPlayer, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	firstPlayer := "Dee"
	stones := []string{
		"101",
		"010",
	}
	expect := false
	runSample(t, n, stones, firstPlayer, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	firstPlayer := "Dum"
	stones := []string{
		"101",
		"010",
	}
	expect := false
	runSample(t, n, stones, firstPlayer, expect)
}
