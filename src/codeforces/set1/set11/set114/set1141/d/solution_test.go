package main

import "testing"

func runSample(t *testing.T, l string, r string, expect int) {
	res := solve(l, r)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := "codeforces"
	r := "dodivthree"
	expect := 5
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l := "abaca?b"
	r := "zabbbcc"
	expect := 5
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l := "bambarbia"
	r := "hellocode"
	expect := 0
	runSample(t, l, r, expect)
}

func TestSample4(t *testing.T) {
	l := "code??????"
	r := "??????test"
	expect := 10
	runSample(t, l, r, expect)
}