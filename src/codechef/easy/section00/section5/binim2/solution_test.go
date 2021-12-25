package main

import "testing"

func runSample(t *testing.T, n int, piles []string, first string, expect string) {
	res := solve(n, piles, first)
	if res != expect {
		t.Errorf("Sample %d %v %s, expect %s, but got %s", n, piles, first, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	piles := []string{"010", "101"}
	first := "Dee"
	runSample(t, n, piles, first, "Dee")
}
