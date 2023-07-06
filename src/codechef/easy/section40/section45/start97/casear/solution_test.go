package main

import "testing"

func runSample(t *testing.T, a, b, c string, expect string) {
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "bd"
	b := "zb"
	c := "dd"
	expect := "bb"
	runSample(t, a, b, c, expect)
}
