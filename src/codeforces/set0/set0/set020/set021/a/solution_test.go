package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "mike___@codeforces.com_.cn"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "mike___@codeforces.com_.cn/xyz"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "john.smith@codeforces.com/contest.icpc/12"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "john_smith@codeforces.com/contest.icpc/12"
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "john_smith@codeforces.com/contest_icpc_12"
	expect := true
	runSample(t, s, expect)
}
