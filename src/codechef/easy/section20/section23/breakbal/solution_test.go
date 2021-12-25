package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := int(solve(S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "()"
	expect := 1
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "(())"
	expect := 2
	runSample(t, S, expect)
}
func TestSample3(t *testing.T) {
	S := "()()"
	expect := 1
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "((()))(())"
	expect := 2
	runSample(t, S, expect)
}
