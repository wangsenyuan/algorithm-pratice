package main

import "testing"

func runSample(t *testing.T, k int, s string, expect string) {
	res := solve(k, s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "whizzy"
	k := 1
	expect := "whizzy"
	runSample(t, k, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aaacdecdebacde"
	k := 2
	expect := "aaaaccdecdebde"
	runSample(t, k, s, expect)
}

func TestSample3(t *testing.T) {
	s := "dbca"
	k := 3
	expect := "abcd"
	runSample(t, k, s, expect)
}

func TestSample4(t *testing.T) {
	s := "acdacc"
	k := 2
	expect := "aacccd"
	runSample(t, k, s, expect)
}

func TestSample5(t *testing.T) {
	s := "z"
	k := 4
	expect := "z"
	runSample(t, k, s, expect)
}
