package main

import "testing"

func runSample(t *testing.T, w string, p int, expect string) {
	res := solve(w, p)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	w := "abca"
	p := 2
	expect := "aa"
	runSample(t, w, p, expect)
}

func TestSample2(t *testing.T) {
	w := "abca"
	p := 6
	expect := "aba"
	runSample(t, w, p, expect)
}
