package main

import "testing"

func runSample(t *testing.T, n int, S string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := "RLRL"
	expect := 0
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	S := "LRRRRL"
	expect := 1
	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	S := "RLLRRRLL"
	expect := 1
	runSample(t, n, S, expect)
}

func TestSample4(t *testing.T) {
	n := 12
	S := "LLLLRRLRRRLL"
	expect := 3
	runSample(t, n, S, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	S := "RRRRR"
	expect := 2
	runSample(t, n, S, expect)
}
