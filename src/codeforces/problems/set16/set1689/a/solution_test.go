package main

import "testing"

func runSample(t *testing.T, a string, b string, k int, expect string) {
	res := solve(a, b, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "aaaaaa"
	b := "bbbb"
	k := 2
	expect := "aabaabaa"
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := "caaca"
	b := "bedededeb"
	k := 3
	expect := "aaabbcc"
	runSample(t, a, b, k, expect)
}


func TestSample3(t *testing.T) {
	a := "noskill"
	b := "wxhtzdy"
	k := 1
	expect := "dihktlwlxnyoz"
	runSample(t, a, b, k, expect)
}
