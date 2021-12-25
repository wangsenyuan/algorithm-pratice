package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "(a+(b*c))"
	expect := "abc*+"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "((a+b)*(z+x))"
	expect := "ab+zx+*"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "((a+t)*((b+(a+c))^(c+d)))"
	expect := "at+bac++cd+^*"
	runSample(t, s, expect)
}
