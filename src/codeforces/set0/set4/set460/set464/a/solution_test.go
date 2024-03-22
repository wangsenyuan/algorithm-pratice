package main

import "testing"

func runSample(t *testing.T, s string, p int, expect string) {
	res := solve(s, p)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cba"
	p := 3
	expect := ""
	runSample(t, s, p, expect)
}
func TestSample2(t *testing.T) {
	s := "cba"
	p := 4
	expect := "cbd"
	runSample(t, s, p, expect)
}

func TestSample3(t *testing.T) {
	s := "abcd"
	p := 4
	expect := "abda"
	runSample(t, s, p, expect)
}

func TestSample4(t *testing.T) {
	s := "cdb"
	p := 4
	expect := "dab"
	runSample(t, s, p, expect)
}

func TestSample5(t *testing.T) {
	s := "yzx"
	p := 26
	expect := "zab"
	runSample(t, s, p, expect)
}

func TestSample6(t *testing.T) {
	s := "abcabcabcab"
	p := 3
	expect := "acbacbacbac"
	runSample(t, s, p, expect)
}
