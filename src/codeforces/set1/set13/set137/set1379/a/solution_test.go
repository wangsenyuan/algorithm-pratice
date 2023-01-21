package main

import "testing"

func runSample(t *testing.T, S string, expect bool) {
	res := solve(len(S), []byte(S))

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "abacaba"
	expect := true
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "???????"
	expect := true
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "aba?abacaba"
	expect := true
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "abacaba?aba"
	expect := true
	runSample(t, S, expect)
}

func TestSample5(t *testing.T) {
	S := "asdf???f???qwer"
	expect := false
	runSample(t, S, expect)
}

func TestSample6(t *testing.T) {
	S := "abacabacaba"
	expect := false
	runSample(t, S, expect)
}
