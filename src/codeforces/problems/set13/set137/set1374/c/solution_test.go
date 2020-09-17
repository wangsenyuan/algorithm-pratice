package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	s := ")("
	expect := 1
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	s := "())()()("
	expect := 1
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	s := ")))((((())"
	expect := 3
	runSample(t, n, s, expect)
}
