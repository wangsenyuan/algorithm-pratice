package main

import "testing"

func runSample(t *testing.T, p []string, w string, lucky string, expect string) {
	res := solve(p, w, lucky)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []string{
		"bers",
		"ucky",
		"elu",
	}
	w := "PetrLoveLuckyNumbers"
	lucky := "t"
	expect := "PetrLovtTttttNumtttt"
	runSample(t, p, w, lucky, expect)
}

func TestSample2(t *testing.T) {
	p := []string{
		"hello",
		"party",
		"abefglghjdhfgj",
		"IVan",
	}
	w := "petrsmatchwin"
	lucky := "a"
	expect := "petrsmatchwin"
	runSample(t, p, w, lucky, expect)
}

func TestSample3(t *testing.T) {
	p := []string{
		"aCa",
		"cba",
	}
	w := "abAcaba"
	lucky := "c"
	expect := "abCacba"
	runSample(t, p, w, lucky, expect)
}
