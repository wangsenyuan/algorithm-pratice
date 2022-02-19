package main

import "testing"

func runSample(t *testing.T, n int, a, b string, expect string) {
	res := solve(n, a, b)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := "aza"
	b := "abz"
	expect := "azabza"
	runSample(t, n, a, b, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := "cccc"
	b := "cccc"
	expect := "cccccccc"
	runSample(t, n, a, b, expect)
}
