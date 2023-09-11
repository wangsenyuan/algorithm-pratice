package main

import "testing"

func runSample(t *testing.T, x string, expect string) {
	res := solve(x)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := "a"
	expect := "b"
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := "ba"
	expect := "ac"
	runSample(t, x, expect)
}

func TestSample3(t *testing.T) {
	x := "codeforces"
	expect := "abcdebfadg"
	runSample(t, x, expect)
}

func TestSample4(t *testing.T) {
	x := "abcdefghijklmnopqrstuvwxyz"
	expect := "bcdefghijklmnopqrstuvwxyza"
	runSample(t, x, expect)
}

func TestSample5(t *testing.T) {
	x := "abcdefghijklmnopqrstuvwxzy"
	expect := "bcdefghijklmnopqrstuvwxyaz"
	runSample(t, x, expect)
}
